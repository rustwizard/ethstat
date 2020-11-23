package eth_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/rustwizard/ethstat/internal/eth"
	"github.com/stretchr/testify/require"
)

func TestDial(t *testing.T) {
	ethc := eth.NewClient(eth.Config{
		URL:        "https://ropsten.infura.io/v3/940d66278ca849f690d6c95a4551c0de",
		RequestTTL: 5 * time.Second,
	})
	err := ethc.Dial()
	require.NoError(t, err)

	ethws := eth.NewClient(eth.Config{
		URL:        "wss://ropsten.infura.io/ws/v3/940d66278ca849f690d6c95a4551c0de",
		RequestTTL: 5 * time.Second,
	})
	err = ethws.Dial()
	require.NoError(t, err)
}

func TestBlockByNumber(t *testing.T) {
	ethws := eth.NewClient(eth.Config{
		URL:        "wss://ropsten.infura.io/ws/v3/940d66278ca849f690d6c95a4551c0de",
		RequestTTL: 5 * time.Second,
	})
	err := ethws.Dial()
	require.NoError(t, err)

	block, err := ethws.BlockByNumber(big.NewInt(1000))
	require.NoError(t, err)
	require.Equal(t, int64(1000), block.Number().Int64())
}

func TestHeaderBlockNum(t *testing.T) {
	ethws := eth.NewClient(eth.Config{
		URL:        "wss://ropsten.infura.io/ws/v3/940d66278ca849f690d6c95a4551c0de",
		RequestTTL: 5 * time.Second,
	})
	err := ethws.Dial()
	require.NoError(t, err)

	headerNum, err := ethws.HeaderBlockNum()
	require.NoError(t, err)
	require.NotNil(t, headerNum)
	t.Log("header_block_num", headerNum.Int64())
}
