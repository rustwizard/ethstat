package cmd

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	migrate "github.com/rubenv/sql-migrate"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

// migrateUpCmd represents the up command
var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "apply migration to DB",
	Run: func(cmd *cobra.Command, args []string) {
		log := zerolog.New(os.Stdout).With().Caller().Logger().With().Str("pkg", "migrate up").Logger()

		dns := fmt.Sprintf("sslmode=%s host=%s port=%s user=%s password='%s' dbname=%s",
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
		n, err = migrate.Exec(db, "postgres", Migrations, migrate.Up)
		if err != nil {
			log.Fatal().Err(err).Msg("apply Migrations")
		}

		log.Info().Msgf("applied %d Migrations", n)
	},
}

func init() {
	migrateCmd.AddCommand(migrateUpCmd)
}
