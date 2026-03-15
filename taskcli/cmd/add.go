/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the too list",
	Long: `Add creates a new task and stores it in the todo list.
	Example:taskcli add "Learn golang"
	This command will add the task "Learn golang" to tasks.json file `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		if len(args) == 0 {
			fmt.Println("Please provide task name")
			return
		}
		tasks := loadTasks()
		task := Task{
			Id:   len(tasks) + 1,
			Name: args[0],
		}
		tasks = append(tasks, task)
		saveTasks(tasks)
		fmt.Println("Task added", task.Name)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
