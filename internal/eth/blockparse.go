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

			for _, tx := range block.Transactions() {
				// get only eth transaction
				// skip ERC-20 token transaction
				if tx.Value().Int64() > 0 {
					msg, err := tx.AsMessage(c.EIPSigner)
					if err != nil {
						c.errCh <- err
						continue
					}

					ethBlock.Txs = append(ethBlock.Txs, Tx{
						BlockNum: block.Number().Int64(),
						ID:       tx.Hash().String(),
						From:     msg.From().String(),
						To:       tx.To().String(),
						Value:    tx.Value(),
					})

					out <- ethBlock
				}
			}
		}
	}()
	return out
}
