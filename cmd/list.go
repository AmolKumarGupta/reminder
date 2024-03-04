package cmd

import (
	"fmt"

	"github.com/AmolKumarGupta/reminder-cli/model"
	"github.com/AmolKumarGupta/reminder-cli/tui/table"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all records",
	Long:  "List all records",
	Run: func(cmd *cobra.Command, args []string) {
		records, err := model.Read()
		if err != nil {
			fmt.Println("Something went wrong while reading")
			return
		}

		header := []string{"Date", "Name", "Description"}

		table.New().
			Header(header).
			Add(records).
			Render()

	},
}
