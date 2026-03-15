/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all tasks in the todo list",
	Long: `List shows all tasks stored in the todo list.

Example:
taskcli list

This command reads tasks from tasks.json and displays them with their Ids`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		tasks := loadTasks()
		for _, t := range tasks {
			fmt.Println(t.Id, "-", t.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
