package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Version of the CLI tool
const Version = "1.0.0"

func main() {
	var rootCmd = &cobra.Command{
		Use:   "cli-tool",
		Short: "A simple CLI tool in Go",
		Long:  `This is a basic CLI tool built using Go and Cobra framework.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Use 'cli-tool help' to see available commands.")
		},
	}

	// Add commands
	rootCmd.AddCommand(greetCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(divyanshCmd)

	// Execute CLI
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
var divyanshCmd = &cobra.Command{
	Use: "divyansh [name]",
	Short: "Divyansh wants to say Thank you",
	Run : func(cmd *cobra.Command , args[] string) {
		fmt.Printf("Hello Divyansh is a boy he is 19 year old boy")
	},
}


// Greet Command
var greetCmd = &cobra.Command{
	Use:   "greet [name]",
	Short: "Greets the user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello, %s! Welcome to the Go CLI tool.\n", args[0])
	},
}

// Version Command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays CLI tool version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("CLI Tool Version: %s\n", Version)
	},
}
