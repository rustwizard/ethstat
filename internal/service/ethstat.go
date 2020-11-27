package service

import (
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

func (e *ETHStat) Run() error {
	e.ethCl.ParseBlocks(e.ethCl.FetchBlocks())
	return nil
}
