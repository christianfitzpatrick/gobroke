// Package connectionhandler implements the state and logic for the broker's server.
package connectionhandler

import (
	"fmt"
	"log"
	"net"
)

// ConnectionHandler implements all server functionality for the broker.
type ConnectionHandler struct {
	host        string
	port        string
	tcpListener *net.TCPListener
}

// NewConnectionHandler instantiates the broker's server.
func NewConnectionHandler(connHost, connPort string) *ConnectionHandler {
	addr := fmt.Sprintf("%s:%s", connHost, connPort)
	laddr, err := net.ResolveTCPAddr("tcp", addr)

	if err != nil {
		return nil
	}

	connListner, listenErr := net.ListenTCP("tcp", laddr)

	if listenErr != nil {
		log.Fatal("Error listening:", listenErr.Error())
	}

	handler := ConnectionHandler{connHost, connPort, connListner}

	return &handler
}

// ListenOnTCP allows a ConnectionHandler to begin accepting TCP connections.
func (connHandler *ConnectionHandler) ListenOnTCP() {
	defer connHandler.tcpListener.Close()

	address := fmt.Sprintf("%s:%s", connHandler.host, connHandler.port)
	log.Println("[TCP] Listening on", address)

	// Start the broker's server and keep looping.
	for {
		// Keep checking for incoming TCP calls.
		conn, connErr := connHandler.tcpListener.AcceptTCP()

		if connErr != nil {
			log.Println("Error listening:", connErr.Error())

			return
		}

		// Instantiate a client connection and manage it as a goroutine.
		client := NewClientConnection(connHandler, conn)
		go client.handleClientConnection()
	}
}
