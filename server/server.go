package server

import (
	"context"
	"fmt"
	"net"

	"github.com/jtyrmn/subreddit-logger-database/database"
	"github.com/jtyrmn/subreddit-logger-database/pb"
	"github.com/jtyrmn/subreddit-logger-database/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
	this file holds the code that implements a grpc server
*/

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
		return &pb.ManyListingsResponse{}, status.Error(codes.Internal, "internal server error")
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
		return &pb.RedditContent{}, status.Error(codes.Internal, "internal server error")
	}

	// not found
	if listing == nil {
		return &pb.RedditContent{}, notFound
	}

	return listing, nil
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
