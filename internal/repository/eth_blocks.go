package repository

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rustwizard/cleargo/db/pg"
)

type EthBlock struct {
	BlockNum int64    `db:"block_num"`
	Txs      []string `db:"txs"`
}

type ETHBlocks interface {
	Put(ctx context.Context, item EthBlock) error
}

type ethBlocks struct {
	db *pg.DB
}

func NewETHBlocks(db *pg.DB) ethBlocks {
	return ethBlocks{db: db}
}

func (r ethBlocks) Put(ctx context.Context, item EthBlock) error {
	_, err := r.db.Pool.Exec(ctx, `INSERT INTO eth_blocks(block_num, txs) VALUES($1, $2)`,
		item.BlockNum, item.Txs)
	if err != nil {
		return errors.Wrap(err, "repository: eth blocks: put item")
	}
	return nil
}
