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

	// Each node can have a maximum of 1 connection for other node
	connections = min(connections, nodes-1)

	maxId := nodes - 1

	for _, n := range nodeArr {
		neighbourIds := getRandomNeighbours(n.ID, 0, maxId, connections)
		fmt.Printf("Node %v: Adding neighbours: %v\n", n.ID, neighbourIds)
		res := getNodesById(nodeArr, neighbourIds...)
		n.Connect(res...)
	}

	for _, n := range nodeArr {
		fmt.Printf("Node %v: Neighbours: %v\n", n.ID, n.Subscribers)
	}

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

func getRandomNeighbours(nodeId base.NodeID, minId int, maxId int, amount int) (n []int) {
	u := newUniqueRand()
	u.Exclude(int(nodeId))

	for i := 0; i < amount; i++ {
		n = append(n, u.Int(minId, maxId))
	}
	return
}

func newUniqueRand() *UniqueRand {
	return &UniqueRand{generated: make(map[int]bool)}
}

type UniqueRand struct {
	generated map[int]bool
}

func (u *UniqueRand) Int(min int, max int) int {
	for {
		i := rand.Intn(max-min+1) + min
		if !u.generated[i] {
			u.generated[i] = true
			return i
		}
	}
}

func (u *UniqueRand) Exclude(num int) {
	u.generated[num] = true
}
