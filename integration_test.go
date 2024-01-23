package main

import (
	"fmt"
	"github.com/vposloncec/go-ssip/base"
	"github.com/vposloncec/go-ssip/orchestration"
	"os"
	"runtime/pprof"
	"testing"
)

func BenchmarkStart(b *testing.B) {
	f, err := os.Create("go-ssip.prof")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Start profiling
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	orchestration.StartRandom(b.N, b.N*10)
	fmt.Printf("Tested with nodes: %v, connections: %v\n", b.N, b.N*10)
}

func TestStartRandom(t *testing.T) {
	nodes := 155000
	connMultiplier := 0.5
	connNum := int(float64(nodes) * float64(connMultiplier))

	orchestration.StartRandom(nodes, connNum)
}

func TestStartFromInput(t *testing.T) {
	nodes := 7
	pairs := []base.ConnectionPair{
		{0, 1},
		{0, 3},
		{1, 2},
		{1, 3},
		{2, 4},
		{3, 6},
		{5, 3},
		{6, 4},
	}

	orchestration.StartFromInput(nodes, pairs)
}
