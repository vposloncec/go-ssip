package main

import (
	"fmt"
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

	fmt.Printf("Testing with nodes: %v, connections: %v\n", b.N, b.N*10)
	orchestration.StartRandom(b.N, b.N*10)
}

func TestStartRandom(t *testing.T) {
	nodes := 5
	connMultiplier := 3

	orchestration.StartRandom(nodes, nodes*connMultiplier)
}

func TestStartFromInput(t *testing.T) {
	nodes := 6
	pairs := []orchestration.ConnectionPair{
		{0, 1},
		{0, 3},
		{1, 3},
		{2, 4},
		{3, 0},
		{3, 6},
	}

	orchestration.StartFromInput(nodes, pairs)
}
