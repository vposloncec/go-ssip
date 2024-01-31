package base

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

type Graph struct {
	log            *zap.SugaredLogger
	nodeLog        *zap.SugaredLogger
	Nodes          []*Node
	Connections    []ConnectionPair
	printAdjacency bool
}

func NewGraph(log *zap.SugaredLogger, nodeNum int, connections []ConnectionPair) *Graph {
	g := Graph{
		Nodes:       make([]*Node, nodeNum),
		Connections: connections,
	}
	g.log = log.Named("Graph")
	g.nodeLog = log.Named("Node")
	g.createNodes()
	g.connectNodes(connections)

	if viper.GetBool("adjacency") || viper.GetBool("verbosity") {
		g.printAdjacency = true
	}

	return &g
}

func (g *Graph) RunPacketReachLoop(packets []*Packet) {
	intervalDuration := viper.GetDuration("reachloop")
	once := viper.GetBool("reachonce")
	if intervalDuration == 0 {
		return
	}

	var elapsedTime time.Duration
	for {
		time.Sleep(intervalDuration)
		elapsedTime += intervalDuration

		g.log.Infof("%v seconds passed, calculating packet reach...", elapsedTime.Seconds())
		for _, p := range packets {
			g.CalcPacketReach(p.ID)
		}
		if once {
			break
		}
	}
}

func (g *Graph) SendPacketsRandomly(packets []*Packet) {
	maxId := len(g.Nodes) - 1
	for _, p := range packets {
		time.Sleep(100 * time.Millisecond)
		g.Nodes[rand.Intn(maxId)].SendPacket(p)
	}
}

func (g *Graph) CalcPacketReach(uuid PacketUUID) {
	var seen int
	for _, node := range g.Nodes {
		if node.AlreadyReceived(uuid) {
			seen++
		}
	}

	percentage := float64(seen) / float64(len(g.Nodes)) * 100
	g.log.Infof("Packet %10v reached %v / %v nodes (%.2f%%)", uuid, seen, len(g.Nodes), percentage)
}

func (g *Graph) PrintAdjacencyList() {
	for _, n := range g.Nodes {
		neighbourIds := make([]int, 0)
		for _, neighbour := range n.Subscribers {
			neighbourIds = append(neighbourIds, int(neighbour.ID))

		}
		if g.printAdjacency {
			g.log.Infow("",
				"Node", n.ID,
				"Neighbours", neighbourIds,
			)
		}
	}
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
	node.Log = g.nodeLog
	g.Nodes[id] = node
}

func (g *Graph) connectNodes(pairs []ConnectionPair) {
	for _, p := range pairs {
		g.Nodes[p[0]].Connect(g.Nodes[p[1]])
	}
}
