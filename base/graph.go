package base

import (
	"fmt"
	"math/rand"
)

type Graph struct {
	Nodes []*Node
}

func NewGraph(nodeNum int, connections []ConnectionPair) *Graph {
	g := Graph{
		Nodes: make([]*Node, nodeNum),
	}
	g.createNodes()
	g.connectNodes(connections)

	return &g
}

func (g *Graph) CalcPacketReach(uuid PacketUUID) {
	var seen int
	for _, node := range g.Nodes {
		if node.AlreadyReceived(uuid) {
			seen++
		}
	}

	percentage := float64(seen) / float64(len(g.Nodes)) * 100
	fmt.Printf("Packet %10v reached %v / %v nodes (%.2f)", uuid, seen, len(g.Nodes), percentage)
}

func GetNodesById(all []*Node, ids ...int) (res []*Node) {
	for _, i := range ids {
		res = append(res, all[i])
	}

	return
}

func (g *Graph) createNodes() {
	for i := range g.Nodes {
		g.createNode(i)
	}
}

// Initialize node with random values
func (g *Graph) createNode(id int) {
	node := NewNode()
	node.ID = NodeID(id)
	node.CpuScore = rand.Intn(20000)
	g.Nodes[id] = node
}

func (g *Graph) connectNodes(pairs []ConnectionPair) {
	for _, p := range pairs {
		g.Nodes[p[0]].Connect(g.Nodes[p[1]])
	}
}
