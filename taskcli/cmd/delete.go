/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task from the todo list",
	Long: `Delete removes a task from todo list using its id.
	Example:taskcli delete 2
	This command will remove the task whose id is 2 from tasks.json`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
		if len(args) == 0 {
			fmt.Println("Please provide valid task")
			return
		}
		tasks = loadTasks()
		id, _ := strconv.Atoi(args[0])
		tasks = loadTasks()
		var newtasks []Task
		for _, t := range tasks {
			if t.Id != id {
				newtasks = append(newtasks, t)
			}
		}
		saveTasks(newtasks)
		fmt.Println("Task deleted")

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

}
