package base

import (
	"fmt"
	"time"
)

type NodeID int

type Node struct {
	ID            NodeID
	Subscribers   []*Node
	PacketHistory map[PacketUUID]*PacketLog
	CpuScore      int
	MessageQueue  chan []*Packet
}

type PacketLog struct {
	recvTime   time.Time
	recvNodeId NodeID
}

func NewNode() *Node {
	return &Node{
		PacketHistory: make(map[PacketUUID]*PacketLog),
		MessageQueue:  make(chan []*Packet),
	}
}

func (n *Node) Connect(nodes ...*Node) {
	for _, neighbour := range nodes {
		n.Subscribers = append(n.Subscribers, neighbour)
		neighbour.AckConn(n)
	}
}

func (n *Node) SendPacket(p *Packet) {
	// fmt.Printf("Node %06d: Sending packet %v\n", n.ID, p.ID)
	n.PacketHistory[p.ID] = &PacketLog{
		recvTime:   time.Now(),
		recvNodeId: n.ID,
	}
	n.sendAll(p)
}

func (n *Node) RecvPacket(callerNode *Node, p *Packet) {

	fmt.Printf("Node %06d: Received packet %v\n", n.ID, p.ID)
	if n.AlreadyReceived(p.ID) {
		fmt.Printf("Node %06d: Packet %v already seen, skipping send\n", n.ID, p.ID)
	} else {
		n.PacketHistory[p.ID] = &PacketLog{
			recvTime:   time.Now(),
			recvNodeId: callerNode.ID,
		}
		n.sendAll(p)
	}
}

func (n *Node) AlreadyReceived(id PacketUUID) bool {
	return n.PacketHistory[id] != nil
}

func (n *Node) AckConn(node *Node) {
	n.Subscribers = append(n.Subscribers, node)
}

func (n *Node) sendAll(p *Packet) {
	for _, neigbour := range n.Subscribers {
		if neigbour.ID == n.PacketHistory[p.ID].recvNodeId {
			continue
		}

		fmt.Printf("Node %06d: Sending packet to node %v, Packet ID: %v\n", n.ID, neigbour.ID, p.ID)
		neigbour.RecvPacket(n, p)
	}
}
