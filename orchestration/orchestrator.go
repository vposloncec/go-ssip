package orchestration

import (
	"fmt"
	"github.com/vposloncec/go-ssip/base"
	"math/rand"
)

func Start(nodes int, connections int) {
	fmt.Println("Hello world from orchestration")
	// Specify the size of the array
	nodeArr := createNodes(nodes)
	maxId := nodes - 1

	connPairs := GenConnectionPairs(0, maxId, connections)
	for _, pair := range connPairs {
		nodeArr[pair[0]].Connect(nodeArr[pair[1]])
	}

	printNodeConnections(nodeArr)

	p := base.NewPacket("asdf")
	nodeArr[0].SendPacket(p)
}

func createNodes(size int) (nodeArr []*base.Node) {
	nodeArr = make([]*base.Node, size)

	// Initialize each element with random values
	for i := range nodeArr {
		node := base.NewNode()
		node.ID = base.NodeID(i)
		node.CpuScore = rand.Intn(20000)
		nodeArr[i] = node
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
		// fmt.Printf("Node %v: Neighbours: %v\n", n.ID, neighbourIds)
	}
}
