package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// migrateDownCmd represents the down command
var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "remove migrates from schema",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate down called")
	},
}

func init() {
	migrateCmd.AddCommand(migrateDownCmd)
}
