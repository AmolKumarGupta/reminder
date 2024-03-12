package cmd

import (
	"github.com/AmolKumarGupta/reminder-cli/config"
	"github.com/spf13/cobra"
)

func global(cmd *cobra.Command) {
	path, _ := cmd.PersistentFlags().GetString("config")
	if path != "" {
		config.Set(path)
	}
}
