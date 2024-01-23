package base

// ConnectionPair represents an edge in a network. It consists of 2 node IDs that are connected together.
// int is used instead of NodeID to provide easier manipulation and better performance
type ConnectionPair [2]int
