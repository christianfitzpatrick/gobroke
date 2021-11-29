package connectionhandler

import (
	"bytes"
	"context"
	"net"
)

// ClientConnection represents the underlying TCP connection for an AMQP connection.
type ClientConnection struct {
	Handler   *ConnectionHandler
	tcpConn   *net.TCPConn
	ctx       context.Context
	ctxCancel context.CancelFunc
}

// NewClientConnection instantiates a new TCP connection for a client service.
func NewClientConnection(handler *ConnectionHandler, conn *net.TCPConn) *ClientConnection {
	return &ClientConnection{
		Handler: handler,
		tcpConn: conn,
	}
}

func (cc *ClientConnection) handleClientConnection() {
	protocolHeaderBuf := make([]byte, 8)
	_, err := cc.tcpConn.Read(protocolHeaderBuf)

	if err != nil {
		// TODO: Close the ClientConnection
		return
	}

	if !verifyProtocolHeader(protocolHeaderBuf) {
		// TODO: Close the ClientConnection
		return
	}

	// TODO: Create a request-scoped context for the ClientConnection.
	cc.ctx, cc.ctxCancel = context.WithCancel(context.Background())

	// TODO: Create a Channel
}


func (cc *ClientConnection) close() {
	// TODO: Implement This
}

// verifyProtocolHeader checks if the Protocol Header frame is valid.
func verifyProtocolHeader(buf []byte) bool {
	// AMQP 0-9-1 Protocol Header.
	protocolHeader := []byte{'A', 'M', 'Q', 'P', 0, 0, 9, 1}

	return bytes.Equal(buf, protocolHeader)
}
