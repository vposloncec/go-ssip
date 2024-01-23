package orchestration

import (
	"fmt"
	"github.com/vposloncec/go-ssip/base"
)

func StartFromInput(nodeNum int, connections []base.ConnectionPair) {
	graph := base.NewGraph(nodeNum, connections)

	printNodeConnections(graph.Nodes)
	p := base.NewPacket("asdf")
	graph.Nodes[0].SendPacket(p)
}

func StartRandom(nodeNum int, connections int) {
	fmt.Println("Hello world from orchestration")

	maxId := nodeNum - 1
	connPairs := GenConnectionPairs(0, maxId, connections)
	graph := base.NewGraph(nodeNum, connPairs)

	printNodeConnections(graph.Nodes)

	p := base.NewPacket("asdf")
	graph.Nodes[0].SendPacket(p)
}

func printNodeConnections(nodeArr []*base.Node) {
	for _, n := range nodeArr {
		neighbourIds := make([]int, 0)
		for _, neighbour := range n.Subscribers {
			neighbourIds = append(neighbourIds, int(neighbour.ID))

		}
		fmt.Printf("Node %v: Neighbours: %v\n", n.ID, neighbourIds)
	}
}
