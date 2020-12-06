package repository

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/rustwizard/cleargo/db/pg"
)

type EthBlockItem struct {
	BlockNum int64    `db:"block_num"`
	Txs      []string `db:"txs"`
}

type TxItem struct {
	ID       string `db:"tx_id"`
	BlockNum int64  `db:"block_num"`
	From     string `db:"from_addr"`
	To       string `db:"to_addr"`
	Value    string `db:"value"`
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

func (r EthBlocks) PutTxs(ctx context.Context, txs []TxItem) error {
	batch := &pgx.Batch{}
	for _, tx := range txs {
		batch.Queue(`INSERT INTO eth_txs(block_num, tx_id, from_addr, to_addr, value) 
							VALUES($1, $2, $3, $4, $5)`, tx.BlockNum, tx.ID, tx.From, tx.To, tx.Value)
	}

	br := r.db.Pool.SendBatch(ctx, batch)
	if _, err := br.Exec(); err != nil {
		return errors.Wrap(err, "repository: eth blocks: put txs")
	}

	return nil
}
