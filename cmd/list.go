/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/kzankpe/task-tracker-cli-go/helper"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the tasks",
	Long: `List all the tasks present in the file regardless the status.
		   The command display the task id, description and status.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		isExist, err := helper.TestTaskFile()
		if err != nil {
			fmt.Println(err)
		}
		if !isExist {
			fmt.Println("No task present at this moment. Please add a task...")
			return
		}
		// Getting element from the current file
		taskList, _ := helper.GetTaskFileContent()
		for _, item := range taskList {

			fmt.Printf("%d %s\n", item.Id, item.Description)

		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
