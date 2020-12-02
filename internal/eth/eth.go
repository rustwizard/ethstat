package eth

import (
	"context"
	"math/big"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

const defaultRequestTTL = 5 * time.Second

type Config struct {
	URL        string        `mapstructure:"ETH_URL" valid:"url,required"`
	RequestTTL time.Duration `valid:"-"`
	FromBlock  *big.Int      `mapstructure:"ETH_FROM_BLOCK" valid:"-"`
}

type Client struct {
	conf      Config
	cl        *ethclient.Client
	chainID   *big.Int
	EIPSigner types.EIP155Signer
	errCh     chan error
}

func NewClient(conf Config) *Client {
	if conf.RequestTTL == 0 {
		conf.RequestTTL = defaultRequestTTL
	}
	return &Client{
		conf:  conf,
		errCh: make(chan error),
	}
}

func (c Config) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return errors.Wrap(err, "validate ethereum client config")
	}

	return nil
}

func (c *Client) Dial() error {
	if err := c.conf.Validate(); err != nil {
		return errors.Wrap(err, "eth client: dial")
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.conf.RequestTTL)
	defer cancel()

	cl, err := ethclient.DialContext(ctx, c.conf.URL)
	if err != nil {
		return errors.Wrap(err, "eth client: dial")
	}

	c.cl = cl

	c.chainID, err = cl.NetworkID(ctx)
	if err != nil {
		return errors.Wrap(err, "eth client: dial")
	}

	if c.conf.FromBlock == nil {
		headerNum, err := c.HeaderBlockNum()
		if err != nil {
			return errors.Wrap(err, "eth client: dial: get header")
		}
		c.conf.FromBlock = headerNum
	}

	c.EIPSigner = types.NewEIP155Signer(c.chainID)

	return nil
}

type Block struct {
	BlockNum int64
	Txs      []Tx
}

type Tx struct {
	ID    string
	From  string
	To    string
	Value *big.Int
}

func (b Block) TxList() []string {
	txList := make([]string, len(b.Txs))
	for i, v := range b.Txs {
		txList[i] = v.ID
	}
	return txList
}
