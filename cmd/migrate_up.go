package cmd

import (
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	migrate "github.com/rubenv/sql-migrate"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

// migrateUpCmd represents the up command
var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "apply migration to DB",
	Run: func(cmd *cobra.Command, args []string) {
		log := zerolog.New(os.Stdout).With().Caller().Logger().With().Str("command", "migrate up").Logger()
		connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&pool_max_conns=%d",
			Conf.DB.User,
			Conf.DB.Password,
			Conf.DB.Host,
			Conf.DB.Port,
			Conf.DB.DatabaseName,
			Conf.DB.SSL,
			Conf.DB.MaxPoolSize,
		)

		config, err := pgx.ParseConfig(connStr)
		if err != nil {
			log.Fatal().Err(err).Msg("parse postgres connection string")
		}

		db := stdlib.OpenDB(*config)

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
