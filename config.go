package pesakit

import (
	"github.com/pesakit/pesakit/pkg/mno"
	"github.com/pesakit/pesakit/pkg/print"
	clix "github.com/urfave/cli/v2"
)

func (c *Client) configCommand() *clix.Command {
	return &clix.Command{
		Name:  "config",
		Usage: "configurations management",
		Subcommands: []*clix.Command{
			c.printConfigCommand(),
		},
	}
}

func (c *Client) printConfigAction() clix.ActionFunc {
	return func(ctx *clix.Context) error {
		printFormat := ctx.String("format")
		mnoChoice := ctx.String("mno")

		pt := print.PayloadTypeFromString(printFormat)
		isVoda := mno.FromString(mnoChoice) == mno.Vodacom
		isTigo := mno.FromString(mnoChoice) == mno.Tigo
		isAirtel := mno.FromString(mnoChoice) == mno.Airtel
		if mnoChoice == "" || mnoChoice == "all" {

			err := print.Out(ctx.Context,"Airtel Config", c.logger, pt, c.airtel.Conf)
			if err != nil {
				return err
			}

			err = print.Out(ctx.Context, "mpesa Config", c.logger, pt, c.mpesa.Conf)
			if err != nil {
				return err
			}

			err = print.Out(ctx.Context,"Tigo Disburse Config",  c.logger, pt, c.tigo.Config.Disburse)
			if err != nil {
				return err
			}

			err = print.Out(ctx.Context,"Tigo pushAction Config",  c.logger, pt, c.tigo.Config.Push)
			if err != nil {
				return err
			}

			return nil
		} else {
			if isVoda {
				err := print.Out(ctx.Context, "mpesa Config", c.logger, pt, c.mpesa.Conf)
				if err != nil {
					return err
				}
				return nil
			}

			if isTigo {
				err := print.Out(ctx.Context,"Tigo pushAction Config",  c.logger, pt, c.tigo.Config.Push)
				if err != nil {
					return err
				}

				err = print.Out(ctx.Context,"Tigo Disburse Config",  c.logger, pt, c.tigo.Config.Disburse)
				if err != nil {
					return err
				}
				return nil
			}

			if isAirtel {
				err := print.Out(ctx.Context,"Airtel Config",  c.logger, pt, c.airtel.Conf)
				if err != nil {
					return err
				}
				return nil
			}

		}

		return nil
	}
}

func (c *Client) printConfigCommand() *clix.Command {
	flags := []clix.Flag{
		&clix.StringFlag{
			Name:  "mno",
			Usage: "mobile money provider (tigo, airtel, vodacom)",
		},
	}

	return &clix.Command{
		Name:  "print",
		Usage: "print configuration set",
		Flags: flags,
		Action: func(ctx *clix.Context) error {
			return c.printConfigAction()(ctx)
		},
	}
}
