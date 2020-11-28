package eth

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rustwizard/ethstat/internal/repository"
)

func (c *Client) ParseBlocks(in <-chan *types.Block) <-chan repository.EthBlockItem {
	out := make(chan repository.EthBlockItem)
	go func() {
		defer close(out)
		for block := range in {
			ethBlock := repository.EthBlockItem{
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
