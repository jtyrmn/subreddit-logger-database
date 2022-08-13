package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

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
type Connection struct {
	connection mongo.Client

	//the listings collection is probably the only collection this file will ever touch
	listings mongo.Collection
}

// call this function to establish a new connection with your mongodb db
func Connect() (*Connection, error) {
	connectionString := util.GetEnv("MONGODB_CONNECTION_STRING")
	databaseName := util.GetEnv("MONGODB_DATABASE_NAME")

	conn, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, errors.New("error connecting to database:\n" + err.Error())
	}

	//"listings" collection
	collection := conn.Database(databaseName).Collection("listings")

	return &Connection{connection: *conn, listings: *collection}, nil
}

// see pb/proto/ListingsDatabase.proto for some context on these functions
// below functions are used in implementations of rpc protocols in the .proto
// above

func (c Connection) ManyListings(limit uint32, skip uint32) ([]*pb.RedditContent, error) {
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
			log.Printf("warning: decoding listing from database failed: %s", err) // TODO: this should be a proper call to the logger, not a regular printf
			continue
		}

		result[result_idx] = listing
		result_idx += 1
	}

	return result[:result_idx], nil
}

/*
FetchListing returns nil,nil in the case that no listing with that ID was
found, and no other errors. Otherwise it will return listing, nil or
nil, error depending on whether a different error occured
*/
func (c Connection) FetchListing(ID string) (*pb.RedditContent, error) {

	response := c.listings.FindOne(context.Background(), bson.D{{"_id", ID}})
	if err := response.Err(); err != nil {
		// listing not found?
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, fmt.Errorf("error querying database: %s", err)
	}

	// DecodeBytes will never return an error if the original operation to
	// create the response object returned no errors
	bsonBytes, _ := response.DecodeBytes()
	listing, err := util.BsonToRedditContent(bsonBytes)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %s", err)
	}

	return listing, nil
}

// returns # of deleted listings
func (c Connection) CullListings(maxAge uint64) (uint32, error) {

	// smallest date of creation before deletion (unix time)
	minTimeOfCreation := uint64(time.Now().Unix()) - maxAge

	response, err := c.listings.DeleteMany(context.Background(), bson.D{{"listing.date", bson.D{{"$lt", minTimeOfCreation}}}})
	if err != nil {
		return 0, fmt.Errorf("error calling database: %s", err)
	}

	return uint32(response.DeletedCount), nil
}

/*
this function outputs it's retrieved listings via the out channel. This is
because this function will be used to in a grpc endpoint that streams it's
result, so it's generally more memory and speed efficient for this function
to stream data out as well.
*/
func (c Connection) RetrieveListings(maxAge uint64, out chan<- *pb.RedditContent, outErr chan<- error) {

	defer close(out)

	minTimeOfCreation := uint64(time.Now().Unix()) - maxAge

	data, err := c.listings.Find(context.Background(), bson.D{{"listing.date", bson.D{{"$gte", minTimeOfCreation}}}})
	if err != nil {
		outErr <- fmt.Errorf("error querying database: %s", err)
		return
	}

	for data.Next(context.Background()) {
		listing, err := util.BsonToRedditContent(data.Current)
		if err != nil {
			// TODO: logging
			log.Printf("warning: decoding listing from database failed: %s", err)
			continue
		}

		out <- listing
	}
	close(out)

	outErr <- nil
}

/*
stream data into in parameter.

use errOut for errors visible to the API consumer
use errOutInteral for internal server errors

only write to one of these channels before exiting. If there are no errors,
write nil to either
*/
func (c Connection) SaveListings(numListings uint32, in <-chan *pb.RedditContent, errOut chan<- error, errOutInternal chan<- error) {
	if numListings == 0 {
		errOut <- nil
		return
	}

	// insert recieved items into a bson-friendly array
	documents := make([]interface{}, numListings)
	documents_idx := 0
	for listing := range in {
		// too many listings sent by client
		if documents_idx >= len(documents) {
			errOut <- errors.New("number of expected listings exceeded")
			return
		}

		// validate listing
		if err := util.IsValidForDatabase(*listing); err != nil {
			log.Printf("warning: listing of ID \"%s\" rejected: %s", listing.Id, err)
			/*
				client never recieves these errors unfortunetly, should probably
				redesign how this works. A bidirectional stream perhaps. TODO
			*/
			continue
		}

		documents[documents_idx] = util.RedditContentToBson(*listing)
		documents_idx += 1
	}

	// in case less documents than expected arrived
	documents = documents[:documents_idx]

	if len(documents) == 0 {
		errOut <- nil
		return
	}

	_, err := c.listings.InsertMany(context.Background(), documents)
	if err != nil && !isDuplicateKeyError(err) { // don't worry about duplicate key errors
		errOutInternal <- fmt.Errorf("error inserting listings into database: %s", err)
		return
	}

	errOut <- nil
}

// duplicate key errors are expected when inserting many listings
func isDuplicateKeyError(err error) bool {
	conv, ok := err.(mongo.BulkWriteException)
	if !ok {
		return false
	}

	for _, writeError := range conv.WriteErrors {
		if writeError.Code == 11000 { //mongodb error code for duplicate key
			return true
		}
	}

	return false
}
