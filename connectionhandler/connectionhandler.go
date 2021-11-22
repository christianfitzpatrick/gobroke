package connectionhandler

import (
	"fmt"
	"log"
	"net"
)

// ConnectionHandler implements all server functionality for the broker.
type ConnectionHandler struct {
	Host        string
	Port        string
	tcpListener net.Listener
}

// NewConnectionHandler instantiates the broker's server.
func NewConnectionHandler(connHost, connPort string) *ConnectionHandler {
	addr := fmt.Sprintf("%s:%s", connHost, connPort)
	connListner, listenErr := net.Listen("tcp", addr)

	if listenErr != nil {
		log.Fatal("Error listening:", listenErr.Error())
	}

	handler := ConnectionHandler{connHost, connPort, connListner}

	return &handler
}
