package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Long:  "get reminder cli version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reminder 0.0.1")
	},
}
