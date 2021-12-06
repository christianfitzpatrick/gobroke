# GoBroke
A poor man's message broker implementing a subset of the AMQP 0-9-1 specification.

The AMQP 0-9-1 spec:
https://www.rabbitmq.com/protocol.html

NOTES:

As of now, the main focus is implementing pub/sub messaging semantics

Currently, routing is based on regular expression matching, but more sophisticated approaches are documented in the RabbitMQ blog:
* https://blog.rabbitmq.com/posts/2010/09/very-fast-and-scalable-topic-routing-part-1
* https://blog.rabbitmq.com/posts/2011/03/very-fast-and-scalable-topic-routing-part-2
* https://blog.rabbitmq.com/posts/2011/03/sender-selected-distribution/
