package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"

	"github.com/jtyrmn/subreddit-logger-database/database"
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
	// checking that input is valid
	if in.Limit > uint32(s.server.MAX_DATABASE_LIMIT) {
		return &pb.ManyListingsResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("limit argument must not exceed %d", s.server.MAX_DATABASE_LIMIT))
	}
	if in.Skip > uint32(s.server.MAX_DATABASE_SKIP) {
		return &pb.ManyListingsResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("skip argument must not exceed %d", s.server.MAX_DATABASE_SKIP))
	}

	// fetch the listings
	listings, err := s.server.databaseInstance.ManyListings(in.Limit, in.Skip)
	if err != nil {
		// TODO: implement logging and record this internal error
		return &pb.ManyListingsResponse{}, status.Error(codes.Internal, INTERNAL_SERVER_ERROR)
	}

	return &pb.ManyListingsResponse{Listings: listings}, err
}

func (s listingsDatabaseServer) FetchListing(ctx context.Context, in *pb.FetchListingRequest) (*pb.RedditContent, error) {
	notFound := status.Error(codes.NotFound, "listing not found")

	/*
		IDs in the database must conform to a specific format, or the database
		object is invalid. If the provided input ID is not of this format, we
		know it will never return anything from the database, therefore we don't
		need to bother with querying the database
	*/
	if !util.IsValidID(in.Id) {
		// should the status code be NotFound or InvalidArgument? Hmmm
		return &pb.RedditContent{}, notFound
	}

	listing, err := s.server.databaseInstance.FetchListing(in.Id)
	if err != nil {
		//TODO logging
		return &pb.RedditContent{}, status.Error(codes.Internal, INTERNAL_SERVER_ERROR)
	}

	// not found
	if listing == nil {
		return &pb.RedditContent{}, notFound
	}

	return listing, nil
}

func (s listingsDatabaseServer) CullListings(ctx context.Context, in *pb.CullListingsRequest) (*pb.CullListingsResponse, error) {
	// safety check
	MIN_CULLING_AGE := util.GetEnvInt("MIN_CULLING_AGE")
	if in.MaxAge < uint64(MIN_CULLING_AGE) {
		return &pb.CullListingsResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("cannot cull listings under minimum culling age of %d seconds", MIN_CULLING_AGE))
	}

	// something else that is also disasterous
	response, err := s.server.databaseInstance.CullListings(in.MaxAge)
	if err != nil {
		//TODO logging
		return &pb.CullListingsResponse{}, status.Error(codes.Internal, INTERNAL_SERVER_ERROR)
	}

	return &pb.CullListingsResponse{NumDeleted: response}, nil
}

func (s listingsDatabaseServer) RetrieveListings(in *pb.RetrieveListingsRequest, stream pb.ListingsDatabase_RetrieveListingsServer) error {
	out := make(chan *pb.RedditContent)
	outErr := make(chan error, 1)

	go s.server.databaseInstance.RetrieveListings(in.MaxAge, out, outErr)

	for listing := range out {
		err := stream.Send(listing)
		if err != nil {
			// TODO: logging
			log.Printf("warning: error sending listing to client: %s", err)
		}
	}

	// return the above RetrieveListings call's output
	if err := <-outErr; err != nil {
		// TODO: logging
		return errors.New(INTERNAL_SERVER_ERROR)
	}

	return nil
}

func (s listingsDatabaseServer) SaveListings(stream pb.ListingsDatabase_SaveListingsServer) error {
	// get the # of listings the client will send
	numListings, err := extractNumListings(stream)
	if err != nil {
		return err
	}

	in := make(chan *pb.RedditContent)
	errChan := make(chan error, 1)
	errChanInternal := make(chan error, 1)
	go s.server.databaseInstance.SaveListings(numListings, in, errChan, errChanInternal)

	for {
		listing, err := stream.Recv()
		if err == io.EOF {
			// finished recieving items, now send response
			break
		}

		if err != nil {
			// TODO: logging
			log.Printf("SaveListings warning: error recieving listing: %s", err)
			break
		}

		in <- listing
	}
	close(in)

	// recieve SaveListings output
	select {
	case err := <-errChan: // regular client-side error
		if err != nil {
			return status.Error(codes.InvalidArgument, err.Error())
		}
	case err := <-errChanInternal: // internal server error
		if err != nil {
			// TODO: logging
			log.Printf("SaveListings error: %s", err)
			return status.Error(codes.Internal, INTERNAL_SERVER_ERROR)
		}
	}

	return stream.SendAndClose(&pb.SaveListingsResponse{})
}

func (s listingsDatabaseServer) UpdateListings(stream pb.ListingsDatabase_UpdateListingsServer) error {
	// get the # of listings the client will send
	numListings, err := extractNumListings(stream)
	if err != nil {
		return err
	}

	in := make(chan *pb.RedditContent)
	errChan := make(chan error, 1)
	errChanInternal := make(chan error, 1)
	go s.server.databaseInstance.UpdateListings(numListings, in, errChan, errChanInternal)

	for {
		listing, err := stream.Recv()
		if err == io.EOF {
			// finished recieving items, now send response
			break
		}

		if err != nil {
			// TODO: logging
			log.Printf("UpdateListings warning: error recieving listing: %s", err)
			break
		}

		in <- listing
	}
	close(in)

	// recieve SaveListings output
	select {
	case err := <-errChan: // regular client-side error
		if err != nil {
			return status.Error(codes.InvalidArgument, err.Error())
		}
	case err := <-errChanInternal: // internal server error
		if err != nil {
			// TODO: logging
			log.Printf("UpdateListings error: %s", err)
			return status.Error(codes.Internal, INTERNAL_SERVER_ERROR)
		}
	}

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
