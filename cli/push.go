package cli

import (
	"github.com/pesakit/pesakit"
	"github.com/urfave/cli/v2"
)

func pushCommand(apiClient *pesakit.Client) *cli.Command {
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
			Aliases: []string{"ref", "id"},
			Usage:   "reference id of the transaction",
		},
	}
	return &cli.Command{
		Name:        "push",
		Usage:       "send push requests to msisdn",
		Description: "send push pay request by specifying  msisdn(phone number), amount, description and reference id",
		Flags:       flags,
		Action:      action(apiClient, pesakit.Push),
	}
}
