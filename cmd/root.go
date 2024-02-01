/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-ssip",
	Short: "A simple gossip algorithm simulation",
	Long: `Go-ssip creates a number of nodes that represent IoT devices
connected together in a P2P network. The nodes can have various attributes
they may or may not be public. Propagation is done using gossip algorithm
with various parameters.

List of currently simulated nodes and edges is available in csv format by default on:
http://localhost:8080/nodes
http://localhost:8080/edges
`,
	Version: "0.1.0",
}

var nodes, connections, packets int
var pReachLoopTime time.Duration

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	viper.SetDefault("author", "Viktor Posloncec <viktor.posloncec@fer.hr>")
	viper.SetDefault("license", "MIT")

	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-ssip.yaml)")

	rootCmd.PersistentFlags().IntVarP(&nodes, "nodes", "n", 100, "number of nodes to spawn")
	rootCmd.PersistentFlags().IntVarP(&connections, "connections", "c", 300, "number of connections each node has to others")
	rootCmd.PersistentFlags().IntVarP(&packets, "packets", "p", 5,
		"number of packets to randomly send through the network, this value should be kept relatively low for large networks")
	rootCmd.PersistentFlags().DurationVarP(&pReachLoopTime, "reachloop", "d", 5*time.Second,
		"Interval in seconds to print the Packet reach calculation. Setting this to 0 will disable Packet reach calculation")

	viper.BindPFlag("nodes", rootCmd.PersistentFlags().Lookup("nodes"))
	viper.BindPFlag("connections", rootCmd.PersistentFlags().Lookup("connections"))
	viper.BindPFlag("packets", rootCmd.PersistentFlags().Lookup("packets"))
	viper.BindPFlag("reachloop", rootCmd.PersistentFlags().Lookup("reachloop"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	adj := rootCmd.PersistentFlags().
		BoolP("adjacency", "l", false, "Print adjacency list")
	verbosity := rootCmd.PersistentFlags().
		BoolP("verbose", "v", false, "Debug output (verbose)")
	reachOnce := rootCmd.PersistentFlags().
		BoolP("reachonce", "o", true, "Run calculate packet reach loop only once")
	port := rootCmd.PersistentFlags().
		Int("port", 8080, "Port to serve csv data on")

	viper.Set("adjacency", adj)
	viper.Set("verbosity", verbosity)
	viper.Set("reachonce", reachOnce)
	viper.Set("port", port)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".go-ssip" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".go-ssip")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
