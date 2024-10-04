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

// markDoneCmd represents the markDone command
var markDoneCmd = &cobra.Command{
	Use:   "markDone",
	Short: "Update task status to Done",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("markDone called")
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
				item.Status = "done"
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
	rootCmd.AddCommand(markDoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markDoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markDoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
