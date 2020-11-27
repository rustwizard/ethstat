package service_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/rustwizard/cleargo/db/pg"
	"github.com/rustwizard/ethstat/internal/eth"
	"github.com/rustwizard/ethstat/internal/repository"
	"github.com/rustwizard/ethstat/internal/service"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ethws := eth.NewClient(eth.Config{
		URL:        "wss://ropsten.infura.io/ws/v3/940d66278ca849f690d6c95a4551c0de",
		RequestTTL: 5 * time.Second,
		FromBlock:  big.NewInt(9151500),
	})
	err := ethws.Dial()
	require.NoError(t, err)

	db := pg.NewDB()
	err = db.Connect(&pg.Config{
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

	svc := service.NewETHStat(service.WithETHClient(ethws), service.WithETHBlockRepository(r))

	for err := range svc.Run(context.Background()) {
		if err != nil {
			t.Error(err)
		}
	}
}
