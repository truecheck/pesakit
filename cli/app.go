/*
 * MIT License
 *
 * Copyright (c) 2021 PESAKIT - MOBILE MONEY TOOLBOX
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

package cli

import (
	"fmt"
	"github.com/pesakit/pesalib/vat"
	"github.com/urfave/cli/v2"
	"io"
	"sort"
	"sync"
	"time"
)

const (
	appName            = "pesakit"
	appShortDesc       = "mobile money toolkit for developers"
	appLongDescription = `
pesakit is a highly configurable commandline tool that comes on handy during testing and development
of systems that integrate with mobile money vendors. With pesakit you can perform a number of tasks
both is sandbox or production environments.

Supported Vendors: Tigo Pesa, Airtel Money and Vodacom MPESA. Hypothetically the tool should work
in countries that the vendors API supports e.g GHANA for MPESA. But the tool has been tested for
Tanzania only

For extensive documentation of usage please visit https://github.com/pesakit/cli/docs

Author:
  - Pius Masengwa Alfred       email: masengwa@pesakit.dev
`
)

type App struct {
	mu            *sync.RWMutex
	writer        io.Writer
	vatCalculator vat.Calculator
	cliApp        *cli.App
}

func (app *App) Run(args []string) error {
	if err := app.cliApp.Run(args); err != nil {
		return fmt.Errorf("could not run application %w", err)
	}

	return nil
}

func (app *App) appWriter() io.Writer {
	app.mu.Lock()
	defer app.mu.Unlock()

	return app.writer
}

func (app *App) AddCommand(commands ...*cli.Command) {
	app.cliApp.Commands = append(app.cliApp.Commands, commands...)
}

func NewApp() *App {
	cliApp := &cli.App{
		Name:                   appName,
		Usage:                  appShortDesc,
		UsageText:              appShortDesc,
		Description:            appLongDescription,
		Commands:               []*cli.Command{},
		Flags:                  []cli.Flag{},
		EnableBashCompletion:   true,
		Compiled:               time.Now(),
		Authors:                nil,
		Copyright:              "MIT License",
		Reader:                 Stdin,
		Writer:                 Stdio,
		ErrWriter:              Stderr,
		ExitErrHandler:         nil,
		Metadata:               nil,
		ExtraInfo:              nil,
		UseShortOptionHandling: true,
	}
	app := &App{
		mu:            &sync.RWMutex{},
		writer:        Stdio,
		vatCalculator: vat.NewCalculator(),
		cliApp:        cliApp,
	}

	writer := app.writer
	app.AddCommand(VatCommand(writer, app.vatCalculator), MnaCommand(writer))
	sort.Sort(cli.FlagsByName(app.cliApp.Flags))
	sort.Sort(cli.CommandsByName(app.cliApp.Commands))

	return app
}
