package message

import (
	"time"

	"github.com/google/uuid"
)

// Message represents application data passed from client to server and from server to client.
type Message struct {
	ID          string
	RoutingKey  string
	PubTime     time.Time
	Payload     []byte
	PayloadSize uint64
}

// NewMessage returns a new message instance.
func NewMessage(routingKey string, payload []byte) *Message {
	// Use a UUID to uniquely identify messages
	uuid := uuid.New().String()

	return &Message{
		ID:         uuid,
		RoutingKey: routingKey,
		PubTime:    time.Now(),
		Payload:    payload,
	}
}
