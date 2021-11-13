package pesakit

import (
	"fmt"
	"github.com/pesakit/cli/io"
	clix "github.com/urfave/cli/v2"
	"os"
)

func cliApp(c *Client) *clix.App {

	desc :=
		`pesakit is a highly configurable commandline tool that comes on handy during testing and
development of systems that integrate with mobile money vendors. With pesakit you can send
C2B (pushpay) requests or B2C (disbursement) requests. You can do this on either production
or staging stage, it just depends on how you configure it. Meaning you can use pesakit in
real production env.

Supported Vendors: Tigo Pesa, Airtel Money and Vodacom MPESA. There is a possibility to use
the tool in countries that the vendors API supports e.g GHANA for MPESA. But the tool has been
tested for Tanzania only.`

	author1 := &clix.Author{
		Name:  "Pius Alfred",
		Email: "me.pius1102@gmail.com",
	}

	verbose := &clix.BoolFlag{
		Name:        "verbose",
		Usage:       "Enable verbose mode",
		Destination: &c.verbose,
	}
	debug := &clix.BoolFlag{
		Name:        "debug",
		Usage:       "Enable debug mode",
		Destination: &c.debug,
	}

	format := &clix.StringFlag{
		Name:  "format",
		Usage: "print format (text, json, yaml)",
	}

	conf := &clix.StringFlag{
		Name:        "conf",
		Usage:       "configuration file path",
		Destination: &c.configFile,
	}

	app := &clix.App{
		Name:                 "pesakit",
		Usage:                "commandline tool to test/interact with Mobile Money API",
		Version:              "1.0.0",
		Description:          desc,
		Commands:             commands(c),
		Flags:                flags(verbose, conf, debug, format),
		EnableBashCompletion: true,
		Before:               beforeActionFunc,
		After:                afterActionFunc,
		CommandNotFound:      onCommand404,
		OnUsageError:         onErrFunc,
		Authors:              authors(author1),
		Copyright:            "MIT Licence, Creative Commons",
		ErrWriter:            os.Stderr,
	}

	return app
}

func beforeActionFunc(context *clix.Context) error {
	return nil
}

func afterActionFunc(context *clix.Context) error {
	return nil
}

func onCommand404(context *clix.Context, s string) {
	_, _ = fmt.Fprintf(io.Stderr, "not found: %s\n", s)
}

func onErrFunc(context *clix.Context, err error, subcommand bool) error {
	_, _ = fmt.Fprintf(io.Stderr, "error: %v\n", err)
	return nil
}

func commands(c *Client) []*clix.Command {
	return appendCommands(
		c.configCommand(),
		c.callbackCommand(),
		c.pushCommand(),
		c.disburseCommand(),
	)
}

func appendCommands(comm ...*clix.Command) []*clix.Command {
	var commands []*clix.Command
	for _, command := range comm {
		commands = append(commands, command)
	}
	return commands
}

func flags(fs ...clix.Flag) []clix.Flag {
	var flgs []clix.Flag
	for _, flg := range fs {
		flgs = append(flgs, flg)
	}
	return flgs
}

func authors(auth ...*clix.Author) []*clix.Author {
	var authors []*clix.Author
	for _, author := range auth {
		authors = append(authors, author)
	}
	return authors
}

func (c *Client) doActionFunc(actionType action) clix.ActionFunc {
	return func(ctx *clix.Context) error {
		phone := ctx.String("phone")
		amount := ctx.Float64("amount")
		desc := ctx.String("description")
		id := ctx.String("reference")

		request := makeRequest(id, amount, phone, desc)

		doResponse, err := c.do(ctx.Context, actionType, request)
		if err != nil {
			return err
		}

		fmt.Printf("response: %v\n", doResponse)
		return nil
	}
}
