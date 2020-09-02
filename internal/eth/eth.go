package eth

import (
	"context"
	"math/big"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

const defaultRequestTTL = 5 * time.Second

type Config struct {
	URL        string        `mapstructure:"URL" valid:"url,required"`
	RequestTTL time.Duration `valid:"-"`
}

type Client struct {
	conf    Config
	cl      *ethclient.Client
	chainID *big.Int
}

func NewClient(conf Config) *Client {
	if conf.RequestTTL == 0 {
		conf.RequestTTL = defaultRequestTTL
	}
	return &Client{
		conf: conf,
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

	return nil
}
