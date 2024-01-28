package orchestration

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/vposloncec/go-ssip/base"
	"github.com/vposloncec/go-ssip/export"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"math/rand"
	"os"
	"time"
)

func StartFromInput(nodeNum int, connections []base.ConnectionPair) {
	log := getLogger()

	graph := base.NewGraph(log, nodeNum, connections)

	graph.PrintAdjacencyList()
	p := base.NewPacket("asdf")
	graph.Nodes[0].SendPacket(p)
	time.Sleep(3 * time.Second)
	graph.CalcPacketReach(p.ID)
	export.EdgesToCSV(graph.Connections).WriteTo(os.Stdout)
	fmt.Println("==============================")
	export.NodesToCSV(graph.Nodes).WriteTo(os.Stdout)
}

func StartRandom(nodeNum int, connections int) *base.Graph {
	log := getLogger()

	maxId := nodeNum - 1
	connPairs := GenConnectionPairs(0, maxId, connections)
	graph := base.NewGraph(log, nodeNum, connPairs)

	graph.PrintAdjacencyList()

	packetn := 40
	packets := make([]*base.Packet, packetn)
	for i := 0; i < packetn; i++ {
		p := base.NewPacket("asdf " + base.NewUUID())
		graph.Nodes[rand.Intn(maxId)].SendPacket(p)
		packets[i] = p
	}

	go func() {
		time.Sleep(10 * time.Second)
		for _, p := range packets {
			graph.CalcPacketReach(p.ID)
		}
	}()

	return graph
}

func getLogger() *zap.SugaredLogger {
	log, _ := zap.NewDevelopment()
	if !viper.GetBool("verbosity") {
		log = log.WithOptions(zap.IncreaseLevel(zapcore.InfoLevel))
	}
	return log.Sugar()
}
