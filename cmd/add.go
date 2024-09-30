/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/kzankpe/task-tracker-cli-go/helper"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add New task to the list",
	Long:  `Add Task to the list. New task will be saved in the json file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		isExist, err := helper.TestTaskFile()
		if err != nil {
			fmt.Println(err)
		}
		if !isExist {
			fmt.Println("File does not exist on the system")
			helper.CreateTaskFile()
		}

		// Getting element from the current file
		taskList, err := helper.GetTaskFileContent()
		size := len(taskList) + 1
		fmt.Println("Opening file to add new task")
		//Adding new task to the file
		newtask := helper.Task{
			Id:          size,
			Description: "Test value",
			Status:      "todo",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		taskList = append(taskList, newtask)
		//Save the information into the task file

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
