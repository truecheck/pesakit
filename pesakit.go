package pesakit

import (
	"github.com/techcraftlabs/pesakit/airtel"
	"github.com/techcraftlabs/pesakit/mpesa"
	"github.com/techcraftlabs/pesakit/pkg/mno"
	"github.com/techcraftlabs/pesakit/tigo"
)

type (
	Client struct {
		AirtelMoney *airtel.Client
		TigoPesa    *tigo.Client
		Mpesa       *mpesa.Client
	}
)

func NewClient(airtelMoney *airtel.Client, tigopesa *tigo.Client, vodaMpesa *mpesa.Client) *Client {
	return &Client{
		AirtelMoney: airtelMoney,
		TigoPesa:    tigopesa,
		Mpesa:       vodaMpesa,
	}
}

func (c *Client) DetermineMNO(phone string) (mno.Operator, string, error) {
	return mno.Get(phone)
}
