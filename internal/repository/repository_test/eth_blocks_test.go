package repository_test

import (
	"context"
	"testing"

	"github.com/rustwizard/cleargo/db/pg"
	"github.com/rustwizard/ethstat/internal/repository"
	"github.com/stretchr/testify/require"
)

func Test_EthBlocks_Put(t *testing.T) {
	db := pg.NewDB()
	err := db.Connect(&pg.Config{
		Host:         "127.0.0.1",
		Port:         5432,
		User:         "postgres",
		Password:     "postgres",
		DatabaseName: "ethstat",
		SSL:          "disable",
		MaxPoolSize:  1000,
	})
	require.NoError(t, err)

	r := repository.NewETHBlocks(db)
	err = r.Put(context.Background(), repository.EthBlock{
		BlockNum: 1000,
		Txs:      []string{"test1", "test2", "test3"},
	})
	require.NoError(t, err)
}
