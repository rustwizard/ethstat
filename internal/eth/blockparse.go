package eth

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var l = log.Logger.With().Str("package", "eth").Logger()

func (c *Client) FetchBlocks(fromBlock int64) <-chan *types.Block {
	out := make(chan *types.Block)

	blockNum := big.NewInt(fromBlock)
	go func() {
		defer close(out)
		for {
			l.Info().Int64("block_num", blockNum.Int64()).Msg("fetch block")
			block, err := c.BlockByNumber(blockNum)
			if err != nil {
				c.errCh <- err
				continue
			}
			out <- block
			blockNum.Add(blockNum, big.NewInt(1))
			headerNum, err := c.HeaderBlockNum()
			if err != nil {
				c.errCh <- err
				continue
			}
			l.Info().Int64("header_block_num", headerNum.Int64()).Msg("fetch block")
			if blockNum.Int64() >= headerNum.Int64() {
				time.Sleep(5 * time.Second)
			}
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

func (c *Client) HeaderBlockNum() (*big.Int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.conf.RequestTTL)
	defer cancel()
	header, err := c.cl.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "get header block")
	}

	return header.Number, nil
}
