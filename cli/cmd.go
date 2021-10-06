package cli

import (
	"fmt"
	"github.com/techcraftlabs/pesakit"
	"github.com/techcraftlabs/pesakit/internal/io"
	clix "github.com/urfave/cli/v2"
	"os"
)

type (
	App struct {
		client *pesakit.Client
		app    *clix.App
	}
)

func commands(client *pesakit.Client) []*clix.Command {
	pushCmd := PushCommand(client).Command()
	disburseCmd := DisburseCommand(client).Command()
	return appendCommands(pushCmd, disburseCmd)
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

func New(httpApiClient *pesakit.Client) *App {

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

	app := &clix.App{
		Name:                 "pesakit",
		Usage:                "commandline tool to test/interact with Mobile Money API",
		UsageText:            "pesakit push|disburse --amount=1000 --reference=\"sending school fees\" --phone=07XXXX9921",
		Version:              "1.0.0",
		Description:          desc,
		Commands:             commands(httpApiClient),
		Flags:                flags(),
		EnableBashCompletion: true,
		Before:               beforeActionFunc,
		After:                afterActionFunc,
		CommandNotFound:      onCommand404,
		OnUsageError:         onErrFunc,
		Authors:              authors(author1),
		Copyright:            "MIT Licence, Creative Commons",
		ErrWriter:            os.Stderr,
	}

	return &App{
		client: httpApiClient,
		app:    app,
	}
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

func (a *App) Run(args []string) error {
	return a.app.Run(args)
}
