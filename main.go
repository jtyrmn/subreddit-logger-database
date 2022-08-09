package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/jtyrmn/subreddit-logger-database/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	//testing ManyListings database functionality

	connection, err := database.Connect()
	if err != nil {
		panic(err)
	}

	// server, err := server.NewServer(connection)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("now listening on %s\n", util.GetEnv("SUBREDDIT_LOGGER_DATABASE_LOCATION"))
	// err = server.Listen()
	// if err != nil {
	// 	panic(err)
	// }

	listings, err := connection.ManyListings(10, 0)
	if err != nil {
		panic(err)
	}

	for idx, listing := range listings {
		t, err := connection.FetchListing(listing.Id)
		if err != nil {
			fmt.Printf("err for %d: %s\n", idx, err)
			continue
		}
		if t == nil {
			fmt.Printf("%d: not found\n", idx)
			continue
		}

		fmt.Printf("%d: %v\n\n", idx, t)
	}
}
