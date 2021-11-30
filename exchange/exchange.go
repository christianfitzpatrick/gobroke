// Package exchange implements state and functionality for an AMQP exchange.
// An exchange is a message routing agent within a virtual host.
// An exchange accepts messages and routing information - principally a routing key -
// and either passes the messages to message queues, or to internal services.
package exchange

import (
	"sync"

	"github.com/christianfitzpatrick/gobroke/binding"
)

// Exchange receives messages from producer applications and
// optionally routes these to message queues within the server.
type Exchange struct {
	knownBindings []*binding.Binding
	mu            sync.Mutex
}

// NewExchange initializes the broker's message exchange.
func NewExchange() *Exchange {
	bindings := make([]*binding.Binding, 0)

	return &Exchange{knownBindings: bindings}
}

// AddBinding appends a non-duplicate binding to the exchange's bindings.
func (e *Exchange) AddBinding(bind *binding.Binding) {
	e.mu.Lock()
	defer e.mu.Unlock()

	// @FromSpec TODO: ignore duplicate bindings
	e.knownBindings = append(e.knownBindings, bind)
}
