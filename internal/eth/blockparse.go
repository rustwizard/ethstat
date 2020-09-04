package eth

import (
	"github.com/ethereum/go-ethereum/core/types"
)

func (c *Client) FetchBlocks() <-chan *types.Block {
	out := make(chan *types.Block)
	go func() {
		defer close(out)
		// TODO: impl logic
	}()
	return out
}
