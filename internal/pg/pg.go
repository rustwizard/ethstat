package pg

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/rs/zerolog"
)

const (
	maxRetry = 10
	ttlRetry = 1 * time.Second
)

type Config struct {
	Host         string `mapstructure:"HOST"`
	Port         int    `mapstructure:"PORT"`
	User         string `mapstructure:"USER"`
	Password     string `mapstructure:"PASSWORD"`
	DatabaseName string `mapstructure:"DB"`
	Schema       string `mapstructure:"SCHEME"`
	SSL          string `mapstructure:"SSL"`
	MaxPoolSize  int    `mapstructure:"POOL_SIZE"`
}

type DB struct {
	Pool *pgxpool.Pool
	log  zerolog.Logger
}

func NewDB(log zerolog.Logger) *DB {
	return &DB{log: log}
}

func (d *DB) Connect(dbc *Config) error {
	args := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&pool_max_conns=%d",
		dbc.User,
		dbc.Password,
		dbc.Host,
		dbc.Port,
		dbc.DatabaseName,
		dbc.SSL,
		dbc.MaxPoolSize,
	)

	poolConfig, err := pgxpool.ParseConfig(args)
	if err != nil {
		d.log.Error().Err(err).Msg("parse config")
		return err
	}

	var db *pgxpool.Pool
	retry := 1
	for retry < maxRetry {
		db, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
		if err != nil {
			d.log.Error().Err(err).Int("retry", retry).
				Dur("second", ttlRetry+(1<<retry)*time.Second).Msg("")
			retry++
			time.Sleep(ttlRetry + (1<<retry)*time.Second)
			continue
		}
		break
	}

	d.Pool = db
	return err
}
