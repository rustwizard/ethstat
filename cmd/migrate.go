package cmd

import (
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A root command to execute SQL migrate scripts",
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
