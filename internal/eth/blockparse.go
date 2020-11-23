package eth

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

func (c *Client) FetchBlocks(fromBlock int64) <-chan *types.Block {
	out := make(chan *types.Block)

	blockNum := big.NewInt(fromBlock)
	go func() {
		defer close(out)
		for {
			time.Sleep(5 * time.Second)
			block, err := c.BlockByNumber(blockNum)
			if err != nil {
				c.errCh <- err
				continue
			}
			out <- block
		}
	}()
	return out
}

func (c *Client) BlockByNumber(blockNum *big.Int) (*types.Block, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.conf.RequestTTL)
	defer cancel()

	block, err := c.cl.BlockByNumber(ctx, blockNum)
	if err != nil {
		return block, errors.Wrap(err, "fetch block: get block by number")
	}

	return block, nil
}
