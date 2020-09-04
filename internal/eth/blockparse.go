package eth

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

func (c *Client) FetchBlocks(fromblock int64) <-chan *types.Block {
	out := make(chan *types.Block)

	blockNum := big.NewInt(fromblock)
	go func() {
		defer close(out)
		for {
			time.Sleep(5 * time.Second)
			block, err := c.blockByNumber(blockNum)
			if err != nil {
				c.errCh <- err
				continue
			}
			out <- block
		}
	}()
	return out
}

func (c *Client) blockByNumber(blocknum *big.Int) (*types.Block, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.conf.RequestTTL)
	defer cancel()

	block, err := c.cl.BlockByNumber(ctx, blocknum)
	if err != nil {
		return block, errors.Wrap(err, "fetch block: get block by number")
	}

	return block, nil
}
