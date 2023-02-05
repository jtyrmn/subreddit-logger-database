package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/jtyrmn/subreddit-logger-database/database"
	"github.com/jtyrmn/subreddit-logger-database/server"
	"github.com/jtyrmn/subreddit-logger-database/util"
)

func main() {
	envPath := ".env"
	if e, exists := os.LookupEnv("ENV_PATH"); exists {
		envPath = e
	}

	err := godotenv.Load(envPath)
	if err != nil {
		panic("cannot load env file: " + err.Error())
	}

	connection, err := database.Connect()
	if err != nil {
		panic("cannot connect to database: " + err.Error())
	}

	server, err := server.NewServer(connection)
	if err != nil {
		panic("cannot create gRPC server: " + err.Error())
	}

	log.Printf("now listening on %s\n\n", util.GetEnv("SUBREDDIT_LOGGER_DATABASE_LOCATION"))
	err = server.Listen()
	if err != nil {
		panic(err)
	}
}
