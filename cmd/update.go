/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/kzankpe/task-tracker-cli-go/helper"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update task using his ID.",
	Long: `This task take two arguments: id and description.
		With the ID, we will retrieve the corresponding task and update
		the description with the new value.
	`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
		fmt.Printf("%s\n", args[0])
		isExist, err := helper.TestTaskFile()
		if err != nil {
			fmt.Println(err)
		}
		if !isExist {
			fmt.Println("No task present at this moment. Please add a task...")
			return
		}
		// Getting element from the current file
		var taskList []helper.Task
		taskList, _ = helper.GetTaskFileContent()
		s, err := strconv.Atoi(args[0])

		for i, item := range taskList {
			if item.Id == s {
				item.Description = args[1]
				item.UpdatedAt = time.Now()
				taskList[i] = item
				break
			}

		}

		updatelist, err := json.MarshalIndent(taskList, "", " ")

		err = helper.SaveTaskfile(updatelist)
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
