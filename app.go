package pesakit

import (
	"fmt"
	"io"
	"sync"

	"github.com/spf13/cobra"
	"github.com/techcraftlabs/airtel"
	xio "github.com/techcraftlabs/base/io"
	"github.com/techcraftlabs/mpesa"
	"github.com/techcraftlabs/tigopesa"
)

const (
	appName            = "pesakit"
	appShortDesc       = "mobile money toolkit for developers"
	appLongDescription = `pesakit is a highly configurable commandline tool that comes on handy during testing and development
of systems that integrate with mobile money vendors. With pesakit you can perform a number of tasks
both is sandbox or production environment such as:
  - encrypting api keys
  - generation session keys
  - sending C2B push requests
  - sending B2C requests to mobile money wallets
  - sending B2B requests to organizations mobile money wallets
  - checking the status of a transaction
  - reversing a transaction
  - direct debit create and payment for MPESA
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
		logger    io.Writer
		debugMode bool
		home      *string
		root      *cobra.Command
		mpesa     *mpesa.Client
		airtel    *airtel.Client
		tigo      *tigopesa.Client
	}
)

func (app *App) getLogger() io.Writer {
	app.mu.RLock()
	defer app.mu.RUnlock()
	return app.logger
}

func (app *App) setLogger(logger io.Writer) {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.logger = logger
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
		logger:    xio.StdErr,
		debugMode: false,
		home:      new(string),
		root:      nil,
		mpesa:     nil,
		airtel:    nil,
		tigo:      nil,
	}

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

func (app *App) setAirtelClient(client *airtel.Client) {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.airtel = client
}

func (app *App) setTigoClient(client *tigopesa.Client) {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.tigo = client
}
