package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// migrateUpCmd represents the up command
var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "apply migrates to  DB schema",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate up called")
	},
}

func init() {
	migrateCmd.AddCommand(migrateUpCmd)
}
