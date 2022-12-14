package util

import (
	"errors"
	"fmt"

	"github.com/jtyrmn/subreddit-logger-database/pb"
	"go.mongodb.org/mongo-driver/bson"
)

/*
middle-man object which bson pulled from database can be directly converted
to (and vice versa), before converting to pb.RedditContent
*/
type BsonStruct struct {
	ID      string `bson:"_id"`
	Listing bsonStructMetadata
	Entries []bsonStructEntry
}

type bsonStructMetadata struct {
	Contenttype string
	Id          string
	Title       string
	Upvotes     uint32
	Comments    uint32
	Date        uint64
	Querydate   uint64
}

type bsonStructEntry struct {
	Upvotes  uint32
	Comments uint32
	Date     uint64
}

// takes in a bson.Raw object and returns a filled pb.RedditContent
// mainly used to convert mongo database compatible data to grpc compatible data
func BsonToRedditContent(bsonBytes bson.Raw) (*pb.RedditContent, error) {

	//decode bson to middle-man object
	var obj BsonStruct
	if err := bson.Unmarshal(bsonBytes, &obj); err != nil {
		return nil, fmt.Errorf("error parsing bson: %s", err)
	}

	// check that important members were recieved
	if obj.ID == "" || obj.Listing.Id == "" {
		return nil, errors.New("listing ID not retrieved")
	}

	//convert middle-man object to proper RedditContent
	//deal with entries array first
	entries := make([]*pb.RedditContent_ListingEntry, len(obj.Entries))
	for idx, entry := range obj.Entries {
		entries[idx] = &pb.RedditContent_ListingEntry{
			Upvotes:     entry.Upvotes,
			Comments:    entry.Comments,
			DateQueried: entry.Date,
		}
	}

	result := pb.RedditContent{
		Id: obj.ID,
		MetaData: &pb.RedditContent_MetaData{
			ContentType: obj.Listing.Contenttype,
			Id:          obj.Listing.Id,
			Title:       obj.Listing.Title,
			Upvotes:     obj.Listing.Upvotes,
			Comments:    obj.Listing.Comments,
			DateCreated: obj.Listing.Date,
			DateQueried: obj.Listing.Querydate,
		},
		Entries: entries,
	}

	return &result, nil
}

/*
converts standard pb.RedditContent to object that can be parsed to
database-compliant bson struct
*/
func RedditContentToBson(r pb.RedditContent) BsonStruct {
	// copy over the entries member
	entries := make([]bsonStructEntry, len(r.Entries))
	for idx, entry := range r.Entries {
		entries[idx] = bsonStructEntry{
			Upvotes:  entry.Upvotes,
			Comments: entry.Comments,
			Date:     entry.DateQueried,
		}
	}
	return BsonStruct{
		ID: r.Id,
		Listing: bsonStructMetadata{
			Contenttype: r.MetaData.ContentType,
			Id:          r.MetaData.Id,
			Title:       r.MetaData.Title,
			Upvotes:     r.MetaData.Upvotes,
			Comments:    r.MetaData.Comments,
			Date:        r.MetaData.DateCreated,
			Querydate:   r.MetaData.DateQueried,
		},
		Entries: entries,
	}
}
