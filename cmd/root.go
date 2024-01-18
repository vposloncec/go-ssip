/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

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
they may or may not be public to other nodes. Propagation is done using gossip algorithm
with various parameters.`,
	Version: "0.1.0",
}

var nodes int
var connections int

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

	rootCmd.PersistentFlags().IntVarP(&nodes, "nodes", "n", 10, "number of nodes to spawn")
	rootCmd.PersistentFlags().IntVarP(&connections, "connections", "c", 3, "number of connections each node has to others")

	viper.BindPFlag("nodes", rootCmd.PersistentFlags().Lookup("nodes"))
	viper.BindPFlag("connections", rootCmd.PersistentFlags().Lookup("connections"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
