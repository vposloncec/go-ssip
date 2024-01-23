package base

import "sync"

// Connection Encapsulates logic of connection the two Nodes.
// It behaves like a broker (agent) between n:m Nodes (usually 1:1) Nodes (Agent in a Pub/Sub design).
// To send a packet, a node broadcasts the packet to the connection. Connection then notifies the receivers about the
// upcoming packet
type Connection struct {
	// Check if Mutex is really needed
	mu        sync.Mutex
	receivers map[string][]chan string
	quit      chan struct{}
	closed    bool
}

func NewConnection() *Connection {
	return &Connection{
		receivers: make(map[string][]chan string),
		quit:      make(chan struct{}),
	}
}

func (b *Connection) Broadcast(topic string, msg string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return
	}

	for _, ch := range b.receivers[topic] {
		ch <- msg
	}
}
