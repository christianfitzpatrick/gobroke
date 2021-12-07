// Package exchange implements state and functionality for an AMQP exchange.
// An exchange is a message routing agent within a virtual host.
// An exchange accepts messages and routing information - principally a routing key -
// and either passes the messages to message queues, or to internal services.
package exchange

import (
	"sync"

	"github.com/christianfitzpatrick/gobroke/binding"
	"github.com/christianfitzpatrick/gobroke/message"
)

// Exchange receives messages from producer applications and
// optionally routes these to message queues within the server.
// NOTE: As of now, the only exchange type implemented is the topic exchange pattern.
// NOTE: We also implement just the default exchange.
type Exchange struct {
	knownBindings []*binding.Binding
	mu            sync.Mutex
}

// NewExchange initializes the broker's default message exchange.
func NewExchange() *Exchange {
	bindings := make([]*binding.Binding, 0)

	return &Exchange{knownBindings: bindings}
}

// AddBinding appends a non-duplicate binding to the exchange's bindings.
func (e *Exchange) AddBinding(bind *binding.Binding) {
	e.mu.Lock()
	defer e.mu.Unlock()

	// @FromSpec: ignore duplicate bindings
	for _, b := range e.knownBindings {
		if bind.Equal(b) {
			return
		}
	}

	e.knownBindings = append(e.knownBindings, bind)
}

// DropBinding removes a binding from the exchange.
func (e *Exchange) DropBinding(bind *binding.Binding) {
	e.mu.Lock()
	defer e.mu.Unlock()

	for idx, b := range e.knownBindings {
		if bind.Equal(b) {
			e.knownBindings = append(e.knownBindings[:idx], e.knownBindings[idx+1:]...)

			return
		}
	}
}

// GetMatchingQueues identifies the relevant queues based on message's routing key.
func (e *Exchange) GetMatchingQueues(msg *message.Message) []string {
	matches := make([]string, 0)

	for _, b := range e.knownBindings {
		if b.MatchTopicString(msg.RoutingKey) {
			matches = append(matches, b.QueueName)
		}
	}

	return matches
}
