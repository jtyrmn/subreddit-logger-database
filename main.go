package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/jtyrmn/subreddit-logger-database/database"
	"github.com/jtyrmn/subreddit-logger-database/server"
	"github.com/jtyrmn/subreddit-logger-database/util"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	connection, err := database.Connect()
	if err != nil {
		panic(err)
	}

	server, err := server.NewServer(connection)
	if err != nil {
		panic(err)
	}

	log.Printf("now listening on %s\n\n", util.GetEnv("SUBREDDIT_LOGGER_DATABASE_LOCATION"))
	err = server.Listen()
	if err != nil {
		panic(err)
	}
}
