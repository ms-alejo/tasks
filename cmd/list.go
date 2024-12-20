/*
Copyright © 2024 Miguel Alejo
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var allFlag bool // flag to show completed tasks

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Long:  `List all uncompleted tasks or include completed tasks with a flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		listTasks(allFlag)
	},
}

func init() {
	listCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "Include completed tasks")
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listTasks(showAll bool) {
	file, err := loadFile("tasks.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading file:", err)
		return
	}
	defer closeFile(file)

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	if len(records) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	// tabular output
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
	if showAll {
		fmt.Fprintln(writer, "ID\tTask\tCreated\tDone")
	} else {
		fmt.Fprintln(writer, "ID\tTask\tCreated")
	}

	// display
	for _, record := range records {
		id, task, createdAt, isComplete := record[0], record[1], record[2], record[3]
		createdTime, _ := time.Parse(time.RFC3339, createdAt)

		if showAll || isComplete == "false" {
			fmt.Fprintf(
				writer,
				"%s\t%s\t%s\t%s\n",
				id,
				task,
				timediff.TimeDiff(createdTime),
				isComplete,
			)
		}
	}

	writer.Flush()
}
