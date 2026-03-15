/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an existing list",
	Long: `Update modifies the name of an existing tasks using its Id.
	Example:taskcli update 1 "Learn go deeply"
	This command updates the task with Id 1 in tasks.json`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
		if len(args) < 2 {
			fmt.Println("Usage:update <id> <new task>")
			return
		}
		id, _ := strconv.Atoi(args[0])
		tasks = loadTasks()
		for i, t := range tasks {
			if t.Id == id {
				tasks[i].Name = args[1]
			}
		}
		saveTasks(tasks)
		fmt.Println("Task updated")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

}
