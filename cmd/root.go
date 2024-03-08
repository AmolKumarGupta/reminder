package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "reminder",
		Short: "reminder",
		Long:  "reminder cli",
		Run: func(cmd *cobra.Command, args []string) {
			if ok, _ := cmd.Flags().GetBool("version"); ok {
				versionCmd.Run(cmd, args)
				return
			}

			UpComingCmd.Run(cmd, args)
		},
	}
)

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "version")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(AddCmd)
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(UpComingCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
