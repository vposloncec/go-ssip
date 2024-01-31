package orchestration

import (
	"github.com/spf13/viper"
	"github.com/vposloncec/go-ssip/base"
	"github.com/vposloncec/go-ssip/web"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func StartFromInput(nodeNum int, connections []base.ConnectionPair) {
	log := getLogger()

	graph := base.NewGraph(log, nodeNum, connections)

	graph.PrintAdjacencyList()
	p := base.NewPacket("asdf")
	graph.Nodes[0].SendPacket(p)
	time.Sleep(6 * time.Second)
	web.Serve(graph)
	graph.CalcPacketReach(p.ID)
}

func StartRandom(nodeNum int, connections int) *base.Graph {
	log := getLogger()

	maxId := nodeNum - 1
	connPairs := GenConnectionPairs(0, maxId, connections)
	graph := base.NewGraph(log, nodeNum, connPairs)

	graph.PrintAdjacencyList()

	packetn := viper.GetInt("packets")
	packets := make([]*base.Packet, packetn)
	for i := 0; i < len(packets); i++ {
		p := base.NewPacket("random message " + base.NewUUID())
		packets[i] = p
	}
	go graph.SendPacketsRandomly(packets)

	go graph.RunPacketReachLoop(packets)

	return graph
}

func getLogger() *zap.SugaredLogger {
	log, _ := zap.NewDevelopment()
	if !viper.GetBool("verbosity") {
		log = log.WithOptions(zap.IncreaseLevel(zapcore.InfoLevel))
	}
	return log.Sugar()
}
