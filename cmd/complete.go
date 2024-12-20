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

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete <taskid>",
	Short: "Mark a task a complete",
	Long:  `Mark a specific task as complete by providing its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID, _ := strconv.Atoi(args[0])
		completeTask(taskID)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func completeTask(taskID int) {
	file, err := loadFile("tasks.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading file:", err)
		return
	}
	defer closeFile(file)

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	// look for task and mark complete
	found := false
	for i, record := range records {
		id, _ := strconv.Atoi(record[0])
		if id == taskID {
			records[i][3] = "true"
			found = true
			break
		}
	}

	if !found {
		fmt.Fprintln(os.Stderr, "Task not found.")
		return
	}

	// rewrite task file with updates
	file.Truncate(0)
	file.Seek(0, io.SeekStart)
	writer := csv.NewWriter(file)
	writer.WriteAll(records)
	writer.Flush()
	fmt.Println("Task marked as complete!")
}
