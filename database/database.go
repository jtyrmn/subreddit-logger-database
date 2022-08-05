package database

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jtyrmn/subreddit-logger-database/pb"
	"github.com/jtyrmn/subreddit-logger-database/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
	this file deals with the database connection (specifically the mongodb
	cluster as of writing this comment)

	much of this file was copied over from the original subreddit-logger repo,
	as I'm moving database functionality over here
*/

// holds an instance of the database connection
type connection struct {
	connection mongo.Client

	//the listings collection is probably the only collection this file will ever touch
	listings mongo.Collection
}

//call this function to establish a new connection with your mongodb db
func Connect() (*connection, error) {
	connectionString := util.GetEnv("MONGODB_CONNECTION_STRING")
	databaseName := util.GetEnv("MONGODB_DATABASE_NAME")

	conn, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, errors.New("error connecting to database:\n" + err.Error())
	}

	//"listings" collection
	collection := conn.Database(databaseName).Collection("listings")

	return &connection{connection: *conn, listings: *collection}, nil
}

// see pb/proto/ListingsDatabase.proto for some context on these functions
// below functions are used in implementations of rpc protocols in the .proto
// above

func (c connection) ManyListings(limit uint32, skip uint32) ([]*pb.RedditContent, error) {
	if limit == 0 {
		return make([]*pb.RedditContent, 0), nil
	}

	// pull from database
	opts := options.Find().SetSort(bson.D{{"listing.date", -1}}).SetSkip(int64(skip)).SetLimit(int64(limit))
	data, err := c.listings.Find(context.Background(), bson.D{}, opts)
	if err != nil {
		return nil, fmt.Errorf("error querying ManyListings: %s", err)
	}

	result := make([]*pb.RedditContent, limit)
	result_idx := 0

	// add all the recieved listings to result array and return it
	for data.Next(context.Background()) {
		listing, err := util.BsonToRedditContent(data.Current)
		if err != nil {
			log.Printf("warning: decoding listing from database failed: %s", err)
			continue
		}

		result[result_idx] = listing
		result_idx += 1
	}

	return result[:result_idx], nil
}
