package main

import (
	"log"

	"github.com/christianfitzpatrick/gobroke/connectionhandler"
)

func main() {
	server := connectionhandler.NewConnectionHandler("localhost", "5672")
	log.Printf("listening on %s:%s", server.Host, server.Port)
}
