package main

import (
	"fmt"
	"log"
	"math"

	"github.com/joho/godotenv"
	"github.com/jtyrmn/subreddit-logger-database/database"
	"github.com/jtyrmn/subreddit-logger-database/reddit"
)

func main() {
	//this service uses the same .env as subreddit-logger
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	//connect to database
	database, err := database.Connect()
	if err != nil {
		log.Fatal("error connecting to database:\n" + err.Error())
	}

	//make sure it works
	listings := make(reddit.ContentGroup)
	count, err := database.RecieveListings(listings, math.MaxInt64)
	if err != nil {
		log.Fatal("error reading database:\n:" + err.Error())
	}

	fmt.Printf("%d listings recieved\n", count)
	for idx, listing := range listings {
		fmt.Printf("%s: %s\n", idx, listing.Title)
	}
}
