package service

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/rustwizard/ethstat/internal/eth"
	"github.com/rustwizard/ethstat/internal/repository"
)

type ETHStat struct {
	ethCl        *eth.Client
	ethBlockRepo repository.EthBlocks
}

type Option func(*ETHStat)

func NewETHStat(opts ...Option) *ETHStat {
	svc := &ETHStat{}
	for _, opt := range opts {
		opt(svc)
	}
	return svc
}

func WithETHClient(ethCl *eth.Client) Option {
	return func(e *ETHStat) {
		e.ethCl = ethCl
	}
}

func WithETHBlockRepository(ethBlockRepo repository.EthBlocks) Option {
	return func(e *ETHStat) {
		e.ethBlockRepo = ethBlockRepo
	}
}

func (e *ETHStat) Run(ctx context.Context) <-chan error {
	return e.saveToDB(ctx, e.ethCl.ParseBlocks(e.ethCl.FetchBlocks()))
}

func (e *ETHStat) saveToDB(ctx context.Context, in <-chan eth.Block) <-chan error {
	errCh := make(chan error)
	txCh := make(chan []eth.Tx)
	go func() {
		defer close(errCh)
		for {
			for block := range in {
				log.Info().Str("pkg", "ethstat").Interface("block", block).Msg("save block")
				errCh <- e.ethBlockRepo.Put(ctx, repository.EthBlockItem{
					BlockNum: block.BlockNum,
					Txs:      block.TxList(),
				})
				txCh <- block.Txs
			}
		}
	}()

	go func() {
		defer close(txCh)
		for {
			for txs := range txCh {
				log.Info().Str("pkg", "ethstat").
					Interface("txs", txs).Msg("save txns")
				errCh <- e.ethBlockRepo.PutTxs(ctx, e.marshalTxs(txs))
			}
		}
	}()

	return errCh
}

func (e *ETHStat) marshalTxs(txs []eth.Tx) []repository.TxItem {
	txItems := make([]repository.TxItem, len(txs))
	for i, tx := range txs {
		txItems[i] = repository.TxItem{
			ID:       tx.ID,
			BlockNum: tx.BlockNum,
			From:     tx.From,
			To:       tx.To,
			Value:    tx.Value.String(),
		}
	}
	return txItems
}
