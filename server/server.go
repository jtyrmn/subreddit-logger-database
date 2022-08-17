package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"strconv"

	"github.com/jtyrmn/subreddit-logger-database/database"
	"github.com/jtyrmn/subreddit-logger-database/logging"
	"github.com/jtyrmn/subreddit-logger-database/pb"
	"github.com/jtyrmn/subreddit-logger-database/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

/*
	this file holds the code that implements a grpc server
*/

const (
	NUM_LISTINGS_HEADER   = "listings-count" // see SaveListings
	INTERNAL_SERVER_ERROR = "internal server error"

	// strings for logging (to identify what rpc was called)
	MANY_LISTINGS    = "ManyListings"
	FETCH_LISTING    = "FetchListing"
	CULL_LISTING     = "CullListing"
	RETRIEVE_LISTING = "RetrieveListing"
	SAVE_LISTINGS    = "SaveListings"
	UPDATE_LISTINGS  = "UpdateListings"
)

type Server struct {
	// a reference to our listings database
	databaseInstance *database.Connection

	// maximum # of listings requested in a single request
	MAX_DATABASE_LIMIT int
	// the mongodb skip operation is O(n)... can't have this be too big either
	MAX_DATABASE_SKIP int
}

// use this to create a new Server object
func NewServer(database *database.Connection) (Server, error) {
	return Server{
		databaseInstance:   database,
		MAX_DATABASE_LIMIT: util.GetEnvInt("MAX_MANYLISTINGS_LIMIT"),
		MAX_DATABASE_SKIP:  util.GetEnvInt("MAX_MANYLISTINGS_SKIP"),
	}, nil
}

type listingsDatabaseServer struct {
	pb.UnimplementedListingsDatabaseServer

	// our server object created above
	// we can access the database through here
	server Server
}

func (s listingsDatabaseServer) ManyListings(ctx context.Context, in *pb.ManyListingsRequest) (*pb.ManyListingsResponse, error) {
	ls := logging.Info(MANY_LISTINGS, ctx, fmt.Sprintf("limit=%d skip=%d", in.Limit, in.Skip))

	// checking that input is valid
	if in.Limit > uint32(s.server.MAX_DATABASE_LIMIT) {
		logging.ClientError(ls, "max limit exceeded")
		return &pb.ManyListingsResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("limit argument must not exceed %d", s.server.MAX_DATABASE_LIMIT))
	}
	if in.Skip > uint32(s.server.MAX_DATABASE_SKIP) {
		logging.ClientError(ls, "max skip exceeded")
		return &pb.ManyListingsResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("skip argument must not exceed %d", s.server.MAX_DATABASE_SKIP))
	}

	// fetch the listings
	listings, err := s.server.databaseInstance.ManyListings(in.Limit, in.Skip)
	if err != nil {
		logging.InternalError(ls, err)
		return &pb.ManyListingsResponse{}, status.Error(codes.Internal, INTERNAL_SERVER_ERROR)
	}

	logging.InfoTail(ls, fmt.Sprintf("returning %s listings", len(listings)))
	return &pb.ManyListingsResponse{Listings: listings}, err
}

func (s listingsDatabaseServer) FetchListing(ctx context.Context, in *pb.FetchListingRequest) (*pb.RedditContent, error) {
	ls := logging.Info(FETCH_LISTING, ctx, fmt.Sprintf("requesting \"%s\"", in.Id))

	notFound := status.Error(codes.NotFound, "listing not found")

	/*
		IDs in the database must conform to a specific format, or the database
		object is invalid. If the provided input ID is not of this format, we
		know it will never return anything from the database, therefore we don't
		need to bother with querying the database
	*/
	if !util.IsValidID(in.Id) {
		// should the status code be NotFound or InvalidArgument? Hmmm
		logging.ClientError(ls, notFound.Error())
		return &pb.RedditContent{}, notFound
	}

	listing, err := s.server.databaseInstance.FetchListing(in.Id)
	if err != nil {
		logging.InternalError(ls, err)
		return &pb.RedditContent{}, status.Error(codes.Internal, INTERNAL_SERVER_ERROR)
	}

	// not found
	if listing == nil {
		logging.ClientError(ls, notFound.Error())
		return &pb.RedditContent{}, notFound
	}

	logging.InfoTail(ls, "returning listing")
	return listing, nil
}

func (s listingsDatabaseServer) CullListings(ctx context.Context, in *pb.CullListingsRequest) (*pb.CullListingsResponse, error) {
	ls := logging.Info(CULL_LISTING, ctx, fmt.Sprintf("culling listing over %d seconds old", in.MaxAge))
	// safety check
	MIN_CULLING_AGE := util.GetEnvInt("MIN_CULLING_AGE")
	if in.MaxAge < uint64(MIN_CULLING_AGE) {
		logging.ClientError(ls, "max age too small")
		return &pb.CullListingsResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("cannot cull listings under minimum culling age of %d seconds", MIN_CULLING_AGE))
	}

	// something else that is also disasterous
	numDeleted, err := s.server.databaseInstance.CullListings(in.MaxAge)
	if err != nil {
		logging.InternalError(ls, err)
		return &pb.CullListingsResponse{}, status.Error(codes.Internal, INTERNAL_SERVER_ERROR)
	}

	logging.InfoTail(ls, fmt.Sprintf("culled %d listing(s)", numDeleted))
	return &pb.CullListingsResponse{NumDeleted: numDeleted}, nil
}

func (s listingsDatabaseServer) RetrieveListings(in *pb.RetrieveListingsRequest, stream pb.ListingsDatabase_RetrieveListingsServer) error {
	ls := logging.Info(RETRIEVE_LISTING, stream.Context(), fmt.Sprintf("retrieving posts of age <=%d seconds", in.MaxAge))

	out := make(chan *pb.RedditContent)
	outErr := make(chan error, 1)

	go s.server.databaseInstance.RetrieveListings(in.MaxAge, out, outErr)

	sentListings := 0
	for listing := range out {
		err := stream.Send(listing)
		if err != nil {
			logging.ClientWarning(ls, fmt.Sprintf("unable to send listing: %s", err))
		} else {
			sentListings += 1
		}
	}

	// return the above RetrieveListings call's output
	if err := <-outErr; err != nil {
		logging.InternalError(ls, err)
		return status.Error(codes.Internal, INTERNAL_SERVER_ERROR)
	}

	logging.InfoTail(ls, fmt.Sprintf("streamed %d listing(s)", sentListings))
	return nil
}

func (s listingsDatabaseServer) SaveListings(stream pb.ListingsDatabase_SaveListingsServer) error {
	var ls logging.LogStruct

	// get the # of listings the client will send
	numListings, err := extractNumListings(stream)
	if err != nil {
		ls = logging.Info(SAVE_LISTINGS, stream.Context(), "saving n/a listings")
		logging.ClientError(ls, fmt.Sprintf("no %s header", NUM_LISTINGS_HEADER))
		return err
	} else {
		ls = logging.Info(SAVE_LISTINGS, stream.Context(), fmt.Sprintf("saving %d listing(s)", numListings))
	}

	in := make(chan *pb.RedditContent)
	errChan := make(chan error, 1)
	errChanInternal := make(chan error, 1)
	go s.server.databaseInstance.SaveListings(numListings, in, errChan, errChanInternal)

	listingsCount := 0
	for {
		listing, err := stream.Recv()
		if err == io.EOF {
			// finished recieving items, now send response
			break
		}

		if err != nil {
			logging.ClientWarning(ls, fmt.Sprintf("cannot recieve listing: %s", err))
			break
		}

		in <- listing
		listingsCount += 1
	}
	close(in)

	// recieve SaveListings output
	select {
	case err := <-errChan: // regular client-side error
		if err != nil {
			logging.ClientError(ls, fmt.Sprintf("%s", err))
			return status.Error(codes.InvalidArgument, err.Error())
		}
	case err := <-errChanInternal: // internal server error
		if err != nil {
			logging.InternalError(ls, err)
			return status.Error(codes.Internal, INTERNAL_SERVER_ERROR)
		}
	}

	/*
		listings sent into database.SaveListings via the in channel are not
		garanteed to be saved in the database (assuming they can be invalid).
		Therefore I can only say "at most _ listings." I could improve this
		by adding another dedicated channel to database.SaveListings to return
		numInserted but it's already cluttered enough with input channels

		TODO: redesign database.SaveListings input + output params
	*/
	logging.InfoTail(ls, fmt.Sprintf("saved at most %d listings", listingsCount))
	return stream.SendAndClose(&pb.SaveListingsResponse{})
}

func (s listingsDatabaseServer) UpdateListings(stream pb.ListingsDatabase_UpdateListingsServer) error {
	var ls logging.LogStruct

	// get the # of listings the client will send
	numListings, err := extractNumListings(stream)
	if err != nil {
		ls = logging.Info(UPDATE_LISTINGS, stream.Context(), "updating n/a listings")
		logging.ClientError(ls, fmt.Sprintf("no %s header", NUM_LISTINGS_HEADER))
		return err
	} else {
		ls = logging.Info(UPDATE_LISTINGS, stream.Context(), fmt.Sprintf("updating %d listings", numListings))
	}

	in := make(chan *pb.RedditContent)
	errChan := make(chan error, 1)
	errChanInternal := make(chan error, 1)
	go s.server.databaseInstance.UpdateListings(numListings, in, errChan, errChanInternal)

	listingsCount := 0
	for {
		listing, err := stream.Recv()
		if err == io.EOF {
			// finished recieving items, now send response
			break
		}

		if err != nil {
			logging.ClientWarning(ls, fmt.Sprintf("cannot recieve listing: %s", err))
			break
		}

		in <- listing
		listingsCount += 1
	}
	close(in)

	// recieve SaveListings output
	select {
	case err := <-errChan: // regular client-side error
		if err != nil {
			logging.ClientError(ls, fmt.Sprintf("%s", err))
			return status.Error(codes.InvalidArgument, err.Error())
		}
	case err := <-errChanInternal: // internal server error
		if err != nil {
			logging.InternalError(ls, err)
			return status.Error(codes.Internal, INTERNAL_SERVER_ERROR)
		}
	}

	/*
		see the InfoTail call in SaveListings...  same issue here

		TODO: redesign database.UpdateListingsResponse input + output params
	*/
	logging.InfoTail(ls, fmt.Sprintf("updated at most %d listings", listingsCount))
	return stream.SendAndClose(&pb.UpdateListingsResponse{})
}

// for SaveListings + UpdateListings funcs
// attempts to parse NUM_LISTINGS_HEADER header from client's request
func extractNumListings(stream grpc.ServerStream) (uint32, error) {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return 0, status.Error(codes.InvalidArgument, "cannot obtain metadata from client")
	}

	numListings, ok := md[NUM_LISTINGS_HEADER]
	if !ok {
		return 0, status.Error(codes.InvalidArgument, fmt.Sprintf("%s header not found", NUM_LISTINGS_HEADER))
	}
	if len(numListings) != 1 {
		return 0, status.Error(codes.InvalidArgument, fmt.Sprintf("%s must have a single value", NUM_LISTINGS_HEADER))
	}

	parsed, err := strconv.Atoi(numListings[0])
	if err != nil || parsed < 0 {
		return 0, status.Error(codes.InvalidArgument, fmt.Sprintf("%s must be an non-negative integer", NUM_LISTINGS_HEADER))
	}

	return uint32(parsed), nil
}

// call this (blocking) function to listen for requests
func (s Server) Listen() error {
	lis, err := net.Listen("tcp", util.GetEnv("SUBREDDIT_LOGGER_DATABASE_LOCATION"))
	if err != nil {
		return fmt.Errorf("error establishing tcp server: %s", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterListingsDatabaseServer(grpcServer, listingsDatabaseServer{server: s})

	err = grpcServer.Serve(lis)
	if err != nil {
		return fmt.Errorf("error serving grpc: %s", err)
	}
	return nil
}
