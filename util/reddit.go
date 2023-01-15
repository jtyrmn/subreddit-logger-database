package util

import (
	"errors"
	"regexp"

	"github.com/jtyrmn/subreddit-logger-database/pb"
)

// the listings in the database must store IDs in the form of t-_------, or
// regex ^t[1-6]_[a-z0-9]{6}$
func IsValidFullID(ID string) bool {
	result, _ := regexp.MatchString("^t[1-6]_[a-z0-9]{7}$", ID)
	return result
}

// for IDs without the t3_-esque prefix
func IsValidID(ID string) bool {
	result, _ := regexp.MatchString("^[a-z0-9]{7}$", ID)
	return result
}

/*
perform this check on RedditContents before attempting to insert them into
the database

returns nil if valid
*/
func IsValidForDatabase(rc *pb.RedditContent) error {
	if !IsValidFullID(rc.Id) {
		return errors.New("invalid ID")
	}
	if !IsValidID(rc.MetaData.Id) {
		return errors.New("invalid metadata ID")
	}

	if rc.Id != rc.MetaData.ContentType+"_"+rc.MetaData.Id {
		return errors.New("ID and metadata ID are unfamiliar")
	}

	return nil
}
