package base

import "math/rand"

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
