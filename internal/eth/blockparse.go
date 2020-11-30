package eth

import (
	"github.com/ethereum/go-ethereum/core/types"
)

func (c *Client) ParseBlocks(in <-chan *types.Block) <-chan Block {
	out := make(chan Block)
	go func() {
		defer close(out)
		for block := range in {
			ethBlock := Block{
				BlockNum: block.Number().Int64(),
			}
			ethBlock.Txs = make([]string, len(block.Transactions()))
			for i, tx := range block.Transactions() {
				ethBlock.Txs[i] = tx.Hash().String()
			}
			out <- ethBlock
		}
	}()
	return out
}
