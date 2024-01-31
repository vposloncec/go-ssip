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

type PacketWithSender struct {
	Packet *Packet
	Sender *Node
}

type Node struct {
	ID               NodeID
	Log              *zap.SugaredLogger
	MessageQueue     chan *PacketWithSender
	Subscribers      []*Node
	mu               sync.Mutex
	PacketHistory    map[PacketUUID]*PacketLog
	Reliability      ReliabilityLevel
	CpuScore         int
	PackagesReceived int
	PackagesSent     int
	PackagesDropped  int
}

type PacketLog struct {
	recvTime   time.Time
	recvNodeId NodeID
}

func NewNode() *Node {
	n := &Node{
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
	p.Timestamp = time.Now()
	n.PacketHistory[p.ID] = &PacketLog{
		recvTime:   p.Timestamp,
		recvNodeId: n.ID,
	}
	n.sendAll(p)
}

func (n *Node) RecvPacket(callerNode *Node, p *Packet) {
	n.MessageQueue <- &PacketWithSender{
		Sender: callerNode,
		Packet: p}
	n.PackagesReceived++
}

func (n *Node) AlreadyReceived(id PacketUUID) bool {
	return n.PacketHistory[id] != nil
}

func (n *Node) AckConn(node *Node) {
	n.Subscribers = append(n.Subscribers, node)
}

func (n *Node) sendAll(p *Packet) {
	n.mu.Lock()
	defer n.mu.Unlock()
	for _, neigbour := range n.Subscribers {
		if neigbour.ID == n.PacketHistory[p.ID].recvNodeId {
			continue
		}

		if ShouldDropPacket(n.Reliability) {
			n.Log.Infof("Node %06d: Dropping packet send to node %v (reliability %v)", n.ID, neigbour.ID, n.Reliability)
			n.PackagesDropped++
			continue
		}

		n.Log.Debugf("Node %06d: Sending packet to node %v, Packet ID: %v\n", n.ID, neigbour.ID, p.ID)
		n.PackagesSent++
		neigbour.RecvPacket(n, p)
	}
}

func (n *Node) packetListener() {
	for {
		p := <-n.MessageQueue
		packet, sender := p.Packet, p.Sender
		n.Log.Debugf("Node %06d: Received packet %v", n.ID, packet.ID)
		if n.AlreadyReceived(packet.ID) {
			continue
			// n.Log.Debugf("Node %06d: Packet %v already seen, skipping send", n.ID, packet.ID)
		} else {
			n.PacketHistory[packet.ID] = &PacketLog{
				recvTime:   time.Now(),
				recvNodeId: sender.ID,
			}
			n.sendAll(packet)
		}
	}
}

func (n *Node) genCPUScore() {
	minScore := 5000
	relMultiplier := 1000

	n.CpuScore = rand.Intn(20000) - rand.Intn(int(n.Reliability)*relMultiplier+1) + minScore
}
