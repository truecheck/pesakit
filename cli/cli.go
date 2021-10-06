package cli

import (
	"fmt"
	"github.com/techcraftlabs/pesakit"
	"github.com/techcraftlabs/pesakit/internal/io"
	"github.com/techcraftlabs/pesakit/pkg/mno"
	clix "github.com/urfave/cli/v2"
)

var _ Commander = (*Cmd)(nil)

const (
	JSON outFormat = iota + 1
	TEXT
)

const (
	Push RequestType = iota
	Disburse
)

type (
	RequestType int
	Cmd         struct {
		ApiClient   *pesakit.Client
		RequestType RequestType
		Name        string
		Usage       string
		Description string
		Flags       []clix.Flag
		SubCommands []*clix.Command
	}
	outFormat int
	Commander interface {
		Command() *clix.Command
		Before(ctx *clix.Context) error
		After(ctx *clix.Context) error
		Action(ctx *clix.Context) error
		OnError(ctx *clix.Context, err error, isSubcommand bool) error
		PrintOut(payload interface{}, format outFormat) error
	}
)

func (c *Cmd) Command() *clix.Command {
	cmd := &clix.Command{
		Name:        c.Name,
		Usage:       c.Usage,
		Description: c.Description,
		Before: func(ctx *clix.Context) error {
			return c.Before(ctx)
		},
		After: func(ctx *clix.Context) error {
			return c.After(ctx)
		},
		Action: func(ctx *clix.Context) error {
			return c.Action(ctx)
		},
		OnUsageError: func(ctx *clix.Context, err error, isSubcommand bool) error {
			return c.OnError(ctx, err, isSubcommand)
		},
		Subcommands: c.SubCommands,
		Flags:       c.Flags,
	}
	return cmd
}

func (c *Cmd) Before(ctx *clix.Context) error {
	return nil
}

func (c *Cmd) After(ctx *clix.Context) error {
	return nil
}

func (c *Cmd) Action(ctx *clix.Context) error {
	reqType := c.RequestType
	phone := ctx.String("phone")
	operator, s, err := c.ApiClient.DetermineMNO(phone)
	if err != nil {
		return err
	}

	return c.action(ctx, reqType, operator, s)
}

func (c *Cmd) OnError(ctx *clix.Context, err error, isSubcommand bool) error {
	return nil
}

func (c *Cmd) PrintOut(payload interface{}, format outFormat) error {
	str, err := jsonString(payload)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(io.Stderr, str)

	return err
}

func (c *Cmd) action(ctx *clix.Context, requestType RequestType, operator mno.Operator, phone string) error {
	switch requestType {
	case Push:
		return c.pushAction(ctx, operator, phone)

	case Disburse:
		return c.disburseAction(ctx, operator, phone)
	}

	return fmt.Errorf("unrecognized action")
}
