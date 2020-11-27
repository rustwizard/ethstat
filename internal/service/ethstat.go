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

func (e *ETHStat) saveToDB(ctx context.Context, in <-chan repository.EthBlock) <-chan error {
	out := make(chan error)
	go func() {
		defer close(out)
		for {
			for block := range in {
				log.Info().Interface("block", block).Msg("ethstat: save block")
				out <- e.ethBlockRepo.Put(ctx, block)
			}
		}
	}()
	return out
}
