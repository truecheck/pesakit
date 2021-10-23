package cli

import (
	"github.com/pesakit/pesakit"
	cli "github.com/urfave/cli/v2"
)

func disburseCommand(apiClient *pesakit.Client) *cli.Command {
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
		Name:        "disburse",
		Usage:       "send money to phone number from your account",
		Description: "send money to specified msisdn, pesakit automatically detect the mno",
		Flags:       flags,
		Action:      action(apiClient, pesakit.Disburse),
	}
}
