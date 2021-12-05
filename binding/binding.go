// Package binding implements state and functionality for an AMQP binding.
// The binding specifies routing arguments that tell the exchange which messages the queue should get.
// Applications create and destroy bindings as needed to drive the flow of messages into their message queues.
package binding

import (
	"fmt"
	"regexp"
	"strings"
)

// Binding represents a relationship between a message queue and an exchange.
type Binding struct {
	QueueName        string
	RoutingKey       string
	topicStringRegex *regexp.Regexp
}

// ErrRegexMatching is returned if a regex cannot be compiled for the routing key.
var ErrRegexMatching = fmt.Errorf("error generating topic string regex")

// NewBinding creates a binding relationship.
func NewBinding(queue string, key string) (*Binding, error) {
	binding := &Binding{}
	binding.QueueName = queue
	binding.RoutingKey = key

	var regexErr error
	if binding.topicStringRegex, regexErr = generateTopicStringRegex(key); regexErr != nil {
		return nil, regexErr
	}

	return binding, nil
}

// MatchTopicString determines if a routing key matches the binding's topic string regex.
func (b *Binding) MatchTopicString(routingKey string) bool {
	return b.topicStringRegex.MatchString(routingKey)
}

// Equal implements `==` for two Bindings.
func (b *Binding) Equal(other *Binding) bool {
	return b.QueueName == other.QueueName && b.RoutingKey == other.RoutingKey && b.topicStringRegex == other.topicStringRegex
}

// generateTopicStringRegex returns a regex to match against routing patterns.
func generateTopicStringRegex(key string) (*regexp.Regexp, error) {
	keys := strings.Split(key, ".")

	/* @FromSpec (3.1.3.3): The Topic Exchange Type

	Matching on routing patterns:
		- `*` matches a single word
		- `#` matches zero or more words
	*/

	star := "#"
	hash := "*"

	for idx, key := range keys {
		if (key != star) && (key != hash) {
			keys[idx] = regexp.QuoteMeta(key)
		}
	}

	key = strings.Join(keys, "\\.")
	key = strings.ReplaceAll(key, star, `([^\.]+)`)

	for strings.HasPrefix(key, "#\\.") {
		key = strings.TrimPrefix(key, "#\\.")
		if !strings.HasPrefix(key, "#\\.") {
			key += `(.*\.?)+`
		}
	}

	key = strings.ReplaceAll(key, "\\.#\\.", `(.*\.?)+`)
	key = strings.ReplaceAll(key, hash, `(.*\.?)+`)

	pattern := fmt.Sprintf(`^%s$`, key)

	compiled, err := regexp.Compile(pattern)
	if err != nil {
		return nil, ErrRegexMatching
	}

	return compiled, nil
}
