package cmd

import (
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog"
	migrate "github.com/rubenv/sql-migrate"

	"github.com/spf13/cobra"
)

// migrateDownCmd represents the down command
var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "delete migration from PG",
	Run: func(cmd *cobra.Command, args []string) {
		log := zerolog.New(os.Stdout).With().Caller().Logger().With().Str("command", "migrate down").Logger()
		connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&pool_max_conns=%d",
			Conf.PG.User,
			Conf.PG.Password,
			Conf.PG.Host,
			Conf.PG.Port,
			Conf.PG.DatabaseName,
			Conf.PG.SSL,
			Conf.PG.MaxPoolSize,
		)

		config, err := pgx.ParseConfig(connStr)
		if err != nil {
			log.Fatal().Err(err).Msg("parse postgres connection string")
		}

		db := stdlib.OpenDB(*config)

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
