package connectionhandler

// The Channel Lifecycle
// 1. The client opens a new channel (Open).
// 2. The server confirms that the new channel is ready (Open-Ok).
// 3. The client and server use the channel as desired.
// 4. One peer (client or server) closes the channel (Close).
// 5. The other peer hand-shakes the channel close (Close-Ok).

// Channel is a bi-directional stream of communications between two AMQP peers.
// Channels allow for multiplexing a heavyweight TCP/IP connection into several
// lightweight conneections.
type Channel struct{}
