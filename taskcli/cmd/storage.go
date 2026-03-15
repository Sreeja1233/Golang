/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// storageCmd represents the storage command
var filename = "tasks.json"
var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks := loadTasks()
		fmt.Println("tasks", tasks)
		saveTasks(tasks)

	},
}
var tasks []Task

func loadTasks() []Task {

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("No existing tasks found")
		return tasks
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error parsing tasks:", err)
		return []Task{}
	}
	return tasks
}
func saveTasks(tasks []Task) {

	data, _ := json.MarshalIndent(tasks, "", " ")
	os.WriteFile(filename, data, 0644)

}

func init() {
	rootCmd.AddCommand(storageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
