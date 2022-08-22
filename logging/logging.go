package logging

import (
	"context"
	"log"

	"github.com/jtyrmn/subreddit-logger-database/util"
	"google.golang.org/grpc/peer"
)

/*
	this file is for logging requests, errors, etc that happen
*/

/*
should client errors/warnings be ignored?
*/
var logClientEventsFlag *bool = nil

type LogStruct struct {
	protocol string
	source   *peer.Peer
}

/*
	call this function at the beginning of each rpc. It takes in the context and logs the client address
*/
// protocol: name of the specific rpc endpoint (SaveListings, CullListings, etc)
// message: info about the request
func Info(protocol string, ctx context.Context, message string) LogStruct {
	var info string
	client, ok := peer.FromContext(ctx)
	if ok {
		info = client.Addr.String()
	} else {
		info = "<unable to retrieve source info>"
	}

	log.Printf("%s: (started %s) %s\n", info, protocol, message)

	return LogStruct{source: client, protocol: protocol}
}

/*
same as Info() except must be called at end of the rpc using the output
returned from Info()
*/
func InfoTail(ls LogStruct, message string) {
	var info string
	if ls.source != nil {
		info = ls.source.Addr.String()
	} else {
		info = "<unable to retrieve source info>"
	}

	log.Printf("%s: (completed %s) %s\n\n", info, ls.protocol, message)
}

// for errors caused by the client
// protocol: name of the specific rpc endpoint (SaveListings, CullListings, etc)
func ClientError(ls LogStruct, err string) {
	if !logClientEvents() {
		return
	}
	log.Printf("\t\033[1;31mclient error:\033[0m %s: %s\n\n", ls.protocol, err)
}

func ClientWarning(ls LogStruct, err string) {
	if !logClientEvents() {
		return
	}
	log.Printf("\t\033[1;31mclient warning:\033[0m %s: %s\n\n", ls.protocol, err)
}

// for internal server errors
// protocol: name of the specific rpc endpoint (SaveListings, CullListings, etc)
func InternalError(ls LogStruct, err error) {
	log.Printf("\t\033[0;31mserver error:\033[0m %s: %v\n\n", ls.protocol, err)
}

func logClientEvents() bool {
	// if the flag isn't set yet, set it
	if logClientEventsFlag == nil {
		env := util.GetEnvDefault("LOG_CLIENT_EVENTS", "true") == "true"
		logClientEventsFlag = &env
	}

	return *logClientEventsFlag
}
