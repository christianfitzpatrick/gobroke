package main

import (
	"log"

	"github.com/christianfitzpatrick/gobroke/connectionhandler"
)

const (
	// Address components for the broker's server.
	// @FromSpec: Standard IANA port number is 5672.
	host = "localhost"
	port = "5672"
)

func main() {
	server := connectionhandler.NewConnectionHandler(host, port)

	server.ListenOnTCP()
	log.Printf("listening on %s:%s", host, port)
}
