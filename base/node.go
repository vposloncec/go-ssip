package base

import "fmt"

type Node struct {
	ID            int
	Subscribers   []*Node
	PacketHistory map[UUID]bool
	CpuScore      int
}

func NewNode() *Node {
	return &Node{
		PacketHistory: make(map[UUID]bool),
	}
}

func (n Node) Connect(nodes ...*Node) {
	for _, neighbour := range nodes {
		n.Subscribers = append(n.Subscribers, neighbour)
		neighbour.AckConn(&n)
	}
}

func (n *Node) SendPacket(p *Packet) {
	fmt.Printf("Node %06d: Sending packet %v\n", n.ID, p.Identifier)

}

func (n *Node) RecvPacket(p *Packet) {
	fmt.Printf("Node %06d: Received packet %v\n", n.ID, p.Identifier)
	n.sendAll(p)

}

func (n *Node) AckConn(node *Node) {
	n.Subscribers = append(n.Subscribers, node)
}

func (n *Node) sendAll(p *Packet) {
}
