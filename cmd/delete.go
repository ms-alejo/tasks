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

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  `Delete a specific task by providing its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID, _ := strconv.Atoi(args[0])
		deleteTask(taskID)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deleteTask(taskID int) {
	file, err := loadFile("tasks.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading file:", err)
		return
	}
	defer closeFile(file)

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	// filter
	newRecords := [][]string{}
	found := false
	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		if id != taskID {
			newRecords = append(newRecords, record)
		} else {
			found = true
		}
	}

	if !found {
		fmt.Fprintln(os.Stderr, "Task not found.")
		return
	}

	file.Truncate(0)
	file.Seek(0, io.SeekStart)
	writer := csv.NewWriter(file)
	writer.WriteAll(newRecords)
	writer.Flush()
	fmt.Println("Task deleted successfully!")
}
