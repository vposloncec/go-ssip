package base

import (
	"go.uber.org/zap"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type NodeID int

func (n NodeID) String() string {
	return strconv.Itoa(int(n))
}

type Node struct {
	ID               NodeID
	Log              *zap.SugaredLogger
	MessageQueue     chan *PacketWithSender
	Subscribers      []*Node
	mu               sync.Mutex
	PacketHistory    map[PacketUUID]*PacketLog
	Reliability      ReliabilityLevel
	CPUScore         int
	PackagesReceived int
	PackagesSent     int
	PackagesDropped  int
}

type PacketLog struct {
	recvTime   time.Time
	recvNodeId NodeID
}

func NewNode(id NodeID) *Node {
	n := &Node{
		ID:            id,
		MessageQueue:  make(chan *PacketWithSender, 1000),
		PacketHistory: make(map[PacketUUID]*PacketLog),
		Reliability:   NewReliability(),
	}
	n.genCPUScore()
	go n.packetListener()

	return n
}

func (n *Node) Connect(nodes ...*Node) {
	for _, neighbour := range nodes {
		n.Subscribers = append(n.Subscribers, neighbour)
		neighbour.AckConn(n)
	}
}

func (n *Node) SendPacket(p *Packet) {
	// fmt.Printf("Node %06d: Sending packet %v\n", n.ID, p.ID)
	n.mu.Lock()
	p.Timestamp = time.Now()
	n.PacketHistory[p.ID] = &PacketLog{
		recvTime:   p.Timestamp,
		recvNodeId: n.ID,
	}
	n.mu.Unlock()
	n.sendAll(p)
}

func (n *Node) RecvPacket(callerNode *Node, p *Packet) {
	n.MessageQueue <- &PacketWithSender{
		Sender: callerNode,
		Packet: p}
}

func (n *Node) AlreadyReceived(id PacketUUID) bool {
	n.mu.Lock()
	defer n.mu.Unlock()
	return n.PacketHistory[id] != nil
}

func (n *Node) AckConn(node *Node) {
	n.Subscribers = append(n.Subscribers, node)
}

func (n *Node) sendAll(p *Packet) {
	n.mu.Lock()
	defer n.mu.Unlock()
	for _, neighbour := range n.Subscribers {
		if neighbour.ID == n.PacketHistory[p.ID].recvNodeId {
			continue
		}

		if ShouldDropPacket(n.Reliability) {
			n.Log.Debugf("Node %06d: Dropping packet send to node %v (reliability %v)", n.ID, neighbour.ID, n.Reliability)
			n.PackagesDropped++
			continue
		}

		n.Log.Debugf("Node %06d: Sending packet to node %v, Packet ID: %v\n", n.ID, neighbour.ID, p.ID)
		n.PackagesSent++
		neighbour.RecvPacket(n, p)
	}
}

func (n *Node) packetListener() {
	for p := range n.MessageQueue {
		packet, sender := p.Packet, p.Sender
		n.PackagesReceived++
		n.Log.Debugf("Node %06d: Received packet %v", n.ID, packet.ID)
		if n.AlreadyReceived(packet.ID) {
			continue
		} else {
			n.recordPacketReceived(packet.ID, sender.ID)
			n.sendAll(packet)
		}
	}
}

func (n *Node) recordPacketReceived(packetID PacketUUID, senderID NodeID) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.PacketHistory[packetID] = &PacketLog{
		recvTime:   time.Now(),
		recvNodeId: senderID,
	}
}

func (n *Node) genCPUScore() {
	minScore := 8000
	reliabilityMultiplier := 2000

	n.CPUScore = rand.Intn(20000) -
		rand.Intn(int(n.Reliability)*reliabilityMultiplier+1) +
		minScore
}
