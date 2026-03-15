/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "taskcli",
	Short: "A simple command line todo manager",
	Long: `TaskCLI is a simple command-line tool to manage your daily tasks.

You can use it to add, list, update, and delete tasks directly from the terminal.
Tasks are stored locally in a JSON file.

Example commands:
taskcli add "Learn Golang"
taskcli list
taskcli update 1 "Practice CLI project"
taskcli delete 2`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("welcome to my cli tool")

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {

		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
