package cli

import (
	"context"
	"fmt"
	"github.com/techcraftlabs/pesakit"
	"github.com/techcraftlabs/pesakit/airtel"
	"github.com/techcraftlabs/pesakit/mpesa"
	"github.com/techcraftlabs/pesakit/pkg/countries"
	"github.com/techcraftlabs/pesakit/pkg/mno"
	"github.com/techcraftlabs/pesakit/tigo"
	"github.com/urfave/cli/v2"
	"time"
)

func PushCommand(client *pesakit.Client) *Cmd {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "phone",
			Aliases: []string{"p"},
			Usage:   "phone number to send push request",
		},

		&cli.Int64Flag{
			Name:    "amount",
			Aliases: []string{"a"},
			Usage:   "amount for push pay",
		},
		&cli.StringFlag{
			Name:    "description",
			Aliases: []string{"d", "desc"},
			Usage:   "transaction message/description",
		},
		&cli.StringFlag{
			Name:    "reference",
			Aliases: []string{"r", "ref", "id"},
			Usage:   "reference id of the transaction",
		},
	}
	return &Cmd{
		ApiClient:   client,
		RequestType: Push,
		Name:        "push",
		Usage:       "send ussd push request",
		Description: "send ussd push request",
		Flags:       flags,
		SubCommands: nil,
	}
}

func (c *Cmd) pushAction(ctx *cli.Context, operator mno.Operator, phone string) error {

	var (
		id     string
		desc   string
		amount int64
	)

	id = ctx.String("reference")
	desc = ctx.String("description")
	amount = ctx.Int64("amount")
	cont := ctx.Context

	switch operator {

	case mno.Airtel:
		return c.airtelPushAction(cont, id, desc, amount, phone)

	case mno.Tigo:
		return c.tigoPushAction(cont, id, desc, amount, phone)

	case mno.Vodacom:
		return c.mpesaPushAction(cont, id, desc, amount, phone)
	}
	return fmt.Errorf("unsupported mno")
}

func (c *Cmd) airtelPushAction(ctx context.Context, id, description string, amount int64, phone string) error {

	req := airtel.PushPayRequest{
		Reference:          description,
		SubscriberCountry:  countries.TANZANIA,
		SubscriberMsisdn:   phone[3:],
		TransactionAmount:  amount,
		TransactionCountry: countries.TANZANIA,
		TransactionID:      id,
	}

	pushPayResponse, err := c.ApiClient.AirtelMoney.Push(ctx, req)
	if err != nil {
		return err
	}

	_ = c.PrintOut(pushPayResponse, JSON)

	return nil
}

func (c *Cmd) mpesaPushAction(ctx context.Context, id, desc string, amount int64, phone string) error {

	req := mpesa.Request{
		ThirdPartyID: fmt.Sprintf("%d", time.Now().Unix()),
		Reference:    id,
		Amount:       float64(amount),
		MSISDN:       phone,
		Desc:         desc,
	}

	pushPayResponse, err := c.ApiClient.Mpesa.PushAsync(ctx, req)
	if err != nil {
		return err
	}

	_ = c.PrintOut(pushPayResponse, JSON)

	return nil
}

func (c *Cmd) tigoPushAction(ctx context.Context, id, desc string, amount int64, phone string) error {
	billerCode := c.ApiClient.TigoPesa.BillerCode
	req := tigo.PayRequest{
		CustomerMSISDN: phone,
		Amount:         amount,
		Remarks:        desc,
		ReferenceID:    fmt.Sprintf("%s%s", billerCode, id),
	}

	pushPayResponse, err := c.ApiClient.TigoPesa.Push(ctx, req)
	if err != nil {
		return err
	}

	_ = c.PrintOut(pushPayResponse, JSON)

	return nil
}
