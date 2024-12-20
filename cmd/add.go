/*
Copyright © 2024 Miguel Alejo
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <task name / description>",
	Short: "Add a new task",
	Long:  `Add a new task to the task list.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]
		addTask(description)
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

func addTask(description string) {
	file, err := loadFile("tasks.csv") // load task with a lock
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading file:", err)
		return
	}
	defer closeFile(file) // ensure file is closed and unlocked

	// read existing tasks
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	// determine the next task ID
	lastID := 0
	if len(records) > 0 {
		lastID, _ = strconv.Atoi(records[len(records)-1][0])
	}

	// create a new task record
	newTask := []string{
		strconv.Itoa(lastID + 1),
		description,
		time.Now().Format(time.RFC3339),
		"false",
	}

	// append the new task to the file
	file.Seek(0, io.SeekEnd)
	writer := csv.NewWriter(file)
	if err := writer.Write(newTask); err != nil {
		fmt.Fprintln(os.Stderr, "Error writing task:", err)
		return
	}

	writer.Flush()
	fmt.Println("Task added successfully!")
}
