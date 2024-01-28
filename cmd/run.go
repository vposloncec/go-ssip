/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vposloncec/go-ssip/orchestration"
	"github.com/vposloncec/go-ssip/web"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start simulation",
	Run: func(cmd *cobra.Command, args []string) {
		nodes, _ := cmd.Flags().GetInt("nodes")
		connections, _ := cmd.Flags().GetInt("connections")
		graph := orchestration.StartRandom(nodes, connections)
		web.Serve(graph)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
