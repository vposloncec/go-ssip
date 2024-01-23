package orchestration

import (
	"fmt"
	"github.com/vposloncec/go-ssip/base"
	"math/rand"
)

func StartFromInput(nodeNum int, connections []base.ConnectionPair) {
	nodes := createNodes(nodeNum)

	for i := 0; i < len(connections); i++ {
		c1, c2 := connections[i][0], connections[i][1]
		nodes[c1].Connect(nodes[c2])
	}

	printNodeConnections(nodes)
	p := base.NewPacket("asdf")
	nodes[0].SendPacket(p)
}

func StartRandom(nodeNum int, connections int) {
	fmt.Println("Hello world from orchestration")
	// Specify the size of the array
	nodes := createNodes(nodeNum)
	maxId := nodeNum - 1

	connPairs := GenConnectionPairs(0, maxId, connections)
	for _, pair := range connPairs {
		nodes[pair[0]].Connect(nodes[pair[1]])
	}

	printNodeConnections(nodes)

	p := base.NewPacket("asdf")
	nodes[0].SendPacket(p)
}

func createNodes(size int) (nodes []*base.Node) {
	nodes = make([]*base.Node, size)

	// Initialize each element with random values
	for i := range nodes {
		node := base.NewNode()
		node.ID = base.NodeID(i)
		node.CpuScore = rand.Intn(20000)
		nodes[i] = node
	}

	return
}

func getNodesById(all []*base.Node, ids ...int) (res []*base.Node) {
	for _, i := range ids {
		res = append(res, all[i])
	}

	return
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
