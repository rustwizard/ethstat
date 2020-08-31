package cmd

import (
	"github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A root command to execute SQL migrate scripts",
}

var Migrations = &migrate.PackrMigrationSource{
	Box: packr.New("Migrations","../scripts/migrations"),
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
