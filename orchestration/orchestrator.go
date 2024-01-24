package orchestration

import (
	"github.com/spf13/viper"
	"github.com/vposloncec/go-ssip/base"
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
	time.Sleep(3 * time.Second)
	graph.CalcPacketReach(p.ID)
}

func StartRandom(nodeNum int, connections int) {
	log := getLogger()

	maxId := nodeNum - 1
	connPairs := GenConnectionPairs(0, maxId, connections)
	graph := base.NewGraph(log, nodeNum, connPairs)

	graph.PrintAdjacencyList()

	p := base.NewPacket("asdf")
	graph.Nodes[0].SendPacket(p)
	time.Sleep(3 * time.Second)
	graph.CalcPacketReach(p.ID)
}

func getLogger() *zap.SugaredLogger {
	log, _ := zap.NewDevelopment()
	if !viper.GetBool("verbosity") {
		log = log.WithOptions(zap.IncreaseLevel(zapcore.InfoLevel))
	}
	return log.Sugar()
}
