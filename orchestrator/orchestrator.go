package orchestrator

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
)

func Start(cmd *cobra.Command, args []string) {
	fmt.Println("Hello world from orchestrator")
	// Specify the size of the array
	arraySize := 5

	// Create an array of structs
	nodes := make([]Node, arraySize)

	// Initialize each element with random values
	for i := 0; i < arraySize; i++ {
		nodes[i] = Node{
			Value: rand.Intn(100), // Adjust the range as needed
		}
	}
}
