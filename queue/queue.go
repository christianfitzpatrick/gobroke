// Package queue implements a (for now) simplified AMQP message queue.
package queue

import (
	"sync"
	"sync/atomic"

	"github.com/christianfitzpatrick/gobroke/message"
)

// Queue represents an AMQP queue for messages.
type Queue struct {
	name       string
	autoDelete bool   // If queue is deleted when no longer in use.
	size       uint64 // Queue size; update this atomically.
	mu         sync.RWMutex
}

// NewQueue constructs an AMQP queue.
func NewQueue(queueName string, autoDel bool) *Queue {
	queue := &Queue{
		name:       queueName,
		autoDelete: autoDel,
	}

	return queue
}

// SafeQueueSize returns the queue's size safely in terms of concurrency.
func (queue *Queue) SafeQueueSize() uint64 {
	return atomic.LoadUint64(&queue.size)
}

// Enqueue pushes a message to the queue.
func (queue *Queue) Enqueue(msg *message.Message) {
	queue.mu.Lock()
	defer queue.mu.Unlock()

	// SAFETY: Atomic update.
	atomic.AddUint64(&queue.size, 1)
}

// Dequeue pops a message from the queue.
func (queue *Queue) Dequeue() {
}
