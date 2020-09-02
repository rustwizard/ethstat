package eth_test

import (
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
