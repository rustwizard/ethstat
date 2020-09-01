package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	migrate "github.com/rubenv/sql-migrate"

	"github.com/spf13/cobra"
)

// migrateDownCmd represents the down command
var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "delete migration from DB",
	Run: func(cmd *cobra.Command, args []string) {
		log := zerolog.New(os.Stdout).With().Caller().Logger().With().Str("pkg", "migrate down").Logger()

		dns := fmt.Sprintf("sslmode=%s host=%s port=%d user=%s password='%s' dbname=%s",
			Conf.DB.SSL,
			Conf.DB.Host,
			Conf.DB.Port,
			Conf.DB.User,
			Conf.DB.Password,
			Conf.DB.DatabaseName,
		)
		db, err := sql.Open("postgres", dns)
		if err != nil {
			log.Fatal().Err(err).Msg("open db")
		}

		var n int
		n, err = migrate.Exec(db, "postgres", Migrations, migrate.Down)
		if err != nil {
			log.Fatal().Err(err).Msg("delete Migrations")
		}

		log.Info().Msgf("deleted %d Migrations", n)
	},
}

func init() {
	migrateCmd.AddCommand(migrateDownCmd)
}
