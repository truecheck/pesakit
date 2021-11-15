package pesakit

import (
	"github.com/techcraftlabs/base/io"
	"github.com/urfave/cli/v2"
)
import libprint "github.com/pesakit/pesakit/pkg/print"

type docsDetails struct {
	Pesakit  string `json:"pesakit" yaml:"pesakit" text:"pesakit"`
	Airtel   string `json:"airtel" yaml:"airtel" text:"airtel"`
	TigoPesa string `json:"tigopesa" yaml:"tigopesa" text:"tigopesa"`
	Mpesa    string `json:"mpesa" yaml:"mpesa" text:"mpesa"`
}

func (c *Client)docsCommand()*cli.Command{
	dd := &docsDetails{
		Pesakit:  "https://github.com/pesakit/pesakit",
		Airtel:   "https://developers.airtel.africa",
		TigoPesa: "https://www.tigo.co.tz/tigo-pesa-for-developers",
		Mpesa:    "https://openapiportal.m-pesa.com/getting-started/dev",
	}
    return &cli.Command{
        Name: "docs",
        Usage: "show documentation details",
        Action: func(ctx *cli.Context) error {
			format := ctx.String("format")
			pt := libprint.PayloadTypeFromString(format)
			if pt == libprint.TEXT{
				pt = libprint.YAML
			}
			err := libprint.Out(ctx.Context, "DOCUMENTATION SITES",io.Stderr, pt,dd)
			if err != nil {
				return err
			}
			return nil
		},
    }
}
