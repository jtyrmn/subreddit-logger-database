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

	data, err := connection.ManyListings(4, 0)
	if err != nil {
		panic(err)
	}

	for _, listing := range data {
		fmt.Println(listing.MetaData.Title)
	}
}
