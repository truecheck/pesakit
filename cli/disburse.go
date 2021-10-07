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

func DisburseCommand(client *pesakit.Client) *Cmd {
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
			Aliases: []string{"desc"},
			Usage:   "disburse message/description",
		},
		&cli.StringFlag{
			Name:    "id",
			Aliases: []string{"reference", "ref"},
			Usage:   "disburse request id",
		},
	}
	return &Cmd{
		ApiClient:   client,
		RequestType: Disburse,
		Name:        "disburse",
		Usage:       "disburse to customer mobile wallet",
		Description: "disburse to customer mobile wallet",
		Flags:       flags,
		SubCommands: nil,
	}
}

func (c *Cmd) disburseAction(ctx *cli.Context, operator mno.Operator, phone string) error {
	var (
		id          string
		amount      int64
		description string
		cont        context.Context
	)
	id = ctx.String("reference")
	description = ctx.String("description")
	amount = ctx.Int64("amount")
	cont = ctx.Context

	switch operator {
	case mno.Airtel:
		return c.airtelDisburseAction(cont, id, description, amount, phone)

	case mno.Tigo:
		return c.tigoDisburseAction(cont, id, amount, phone)

	case mno.Vodacom:
		return c.mpesaDisburseAction(cont, id, description, amount, phone)
	}
	return fmt.Errorf("unsupported mno")
}

func (c *Cmd) airtelDisburseAction(ctx context.Context, id, description string, amount int64, phone string) error {
	req := airtel.DisburseRequest{
		Reference:            description,
		MSISDN:               phone[3:],
		Amount:               amount,
		CountryOfTransaction: countries.TANZANIA,
		ID:                   id,
	}
	disburseResponse, err := c.ApiClient.AirtelMoney.Disburse(ctx, req)
	if err != nil {
		return err
	}
	err = c.PrintOut(disburseResponse, JSON)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cmd) tigoDisburseAction(ctx context.Context, id string, amount int64, phone string) error {
	req := tigo.Request{
		ReferenceID: id,
		MSISDN:      phone,
		Amount:      float64(amount),
	}

	disburseResponse, err := c.ApiClient.TigoPesa.Disburse(ctx, req)
	if err != nil {
		return err
	}
	err = c.PrintOut(disburseResponse, JSON)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cmd) mpesaDisburseAction(ctx context.Context, id, description string, amount int64, phone string) error {
	req := mpesa.Request{
		ThirdPartyID: fmt.Sprintf("%d", time.Now().UnixNano()),
		Reference:    id,
		Amount:       float64(amount),
		MSISDN:       phone,
		Desc:         description,
	}

	disburseResponse, err := c.ApiClient.Mpesa.Disburse(ctx, req)
	if err != nil {
		return err
	}
	err = c.PrintOut(disburseResponse, JSON)
	if err != nil {
		return err
	}
	return nil
}
