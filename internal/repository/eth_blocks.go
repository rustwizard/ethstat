package repository

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rustwizard/cleargo/db/pg"
)

type EthBlockItem struct {
	BlockNum int64    `db:"block_num"`
	Txs      []string `db:"txs"`
}

type EthBlocks struct {
	db *pg.DB
}

func NewETHBlocks(db *pg.DB) EthBlocks {
	return EthBlocks{db: db}
}

func (r EthBlocks) Put(ctx context.Context, item EthBlockItem) error {
	_, err := r.db.Pool.Exec(ctx, `INSERT INTO eth_blocks(block_num, txs) VALUES($1, $2)`,
		item.BlockNum, item.Txs)
	if err != nil {
		return errors.Wrap(err, "repository: eth blocks: put item")
	}
	return nil
}
