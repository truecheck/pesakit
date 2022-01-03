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

package pesakit

import (
	"fmt"
	"io"
	"log"
	"sync"

	xlog "github.com/pesakit/pesakit/log"
	"github.com/spf13/cobra"
	"github.com/techcraftlabs/airtel"
	"github.com/techcraftlabs/mpesa"
	"github.com/techcraftlabs/tigopesa"
)

const (
	appName            = "pesakit"
	appShortDesc       = "mobile money toolkit for developers"
	appLongDescription = `
pesakit is a highly configurable commandline tool that comes on handy during testing and development
of systems that integrate with mobile money vendors. With pesakit you can perform a number of tasks
both is sandbox or production environment such as:
  1. encrypting api keys
  2. generation session keys
  3. sending C2B push requests
  4. sending B2C requests to mobile money wallets
  5. sending B2B requests to organizations mobile money wallets
  6. checking the status of a transaction
  7. reversing a transaction
  8. direct debit create and payment for MPESA
Supported Vendors: Tigo Pesa, Airtel Money and Vodacom MPESA. Hypothetically the tool should work
in countries that the vendors API supports e.g GHANA for MPESA. But the tool has been tested for
Tanzania only

For extensive documentation of usage please visit https://github.com/pesakit/cli/docs

Author:
  - Pius Masengwa Alfred       email: masengwa@pesakit.dev
`
)

type (
	App struct {
		mu        *sync.RWMutex
		writer    io.Writer
		debugMode bool
		home      *string
		root      *cobra.Command
		mpesa     *mpesa.Client
		airtel    *airtel.Client
		tigo      *tigopesa.Client
		logger    *log.Logger
	}
)

func (app *App) Logger() *log.Logger {
	app.mu.RLock()
	defer app.mu.RUnlock()
	return app.logger
}

func (app *App) getWriter() io.Writer {
	app.mu.RLock()
	defer app.mu.RUnlock()

	return app.writer
}

func (app *App) setWriter(logger io.Writer) {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.writer = logger
}

func (app *App) setDebugMode(debugMode bool) {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.debugMode = debugMode
}

func (app *App) getDebugMode() bool {
	app.mu.RLock()
	defer app.mu.RUnlock()
	return app.debugMode
}

func New() *App { //nolint:ireturn
	app := &App{
		mu:        &sync.RWMutex{},
		writer:    xlog.StdErr,
		debugMode: false,
		home:      new(string),
		mpesa:     nil,
		airtel:    nil,
		tigo:      nil,
	}
	app.logger = xlog.New()
	app.createRootCommand()
	return app
}

func (app *App) Execute() error {
	if err := app.root.Execute(); err != nil {
		return fmt.Errorf("failed to execute: %w", err)
	}

	return nil
}

func (app *App) setHomeDir(homeDir string) {
	app.mu.Lock()
	defer app.mu.Unlock()
	*app.home = homeDir
}

func (app *App) getHome() string {
	app.mu.RLock()
	defer app.mu.RUnlock()
	return *app.home
}

func (app *App) setMpesaClient(client *mpesa.Client) {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.mpesa = client
}

func (app *App) MpesaClient() *mpesa.Client {
	app.mu.RLock()
	defer app.mu.RUnlock()
	return app.mpesa
}

func (app *App) setAirtelClient(client *airtel.Client) {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.airtel = client
}

func (app *App) AirtelClient() *airtel.Client {
	app.mu.RLock()
	defer app.mu.RUnlock()
	return app.airtel
}

func (app *App) setTigoClient(client *tigopesa.Client) {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.tigo = client
}

func (app *App) TigoClient() *tigopesa.Client {
	app.mu.RLock()
	defer app.mu.RUnlock()
	return app.tigo
}
