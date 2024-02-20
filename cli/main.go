package main

import (
	"fmt"
	"os"

	"github.com/StreamMQ/StreamMQ/cli/cmd"
	"github.com/spf13/cobra"

)

var rootCmd = &cobra.Command{
	Use:   "fluffy-cli",
	Short: "A CLI tool for interacting with MQTT",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to MQTT CLI!")
	},
}

func main() {
	rootCmd.AddCommand(cmd.ConnectCmd)
	rootCmd.AddCommand(cmd.PublishCmd)
	rootCmd.AddCommand(cmd.SubscribeCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
