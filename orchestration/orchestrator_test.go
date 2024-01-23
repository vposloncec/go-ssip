package orchestration

import (
	"fmt"
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
	Start(b.N, b.N*10)
}

func TestStart(t *testing.T) {
	nodes := 100000
	connMultiplier := 20

	Start(nodes, nodes*connMultiplier)
}
