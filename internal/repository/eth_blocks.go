package repository

import "context"

type EthBlock struct {
	BlockNum int64    `db:"block_num"`
	Txs      []string `db:"txs"`
}

type ETHBlocks interface {
	Put(ctx context.Context, item EthBlock) error
}
