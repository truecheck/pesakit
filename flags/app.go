package flags

import (
	"github.com/pesakit/pesakit/config"
	"github.com/spf13/cobra"
)

const (
	DebugName          = "debug"
	debugModeUsage     = "debug mode"
	HomeDirectoryName  = "home"
	homeDirectoryUsage = "the home directory of the pesakit application"
	ConfigName         = "config"
	configFileUsage    = "config file"
)

func SetAppFlags(cmd *cobra.Command, app *config.App) {
	if app == nil {
		app = config.DefaultAppConf()
	}
	cmd.PersistentFlags().StringVar(&app.Home, HomeDirectoryName, app.Home, homeDirectoryUsage)
	cmd.PersistentFlags().StringVar(&app.Config, ConfigName, app.Config, configFileUsage)
	cmd.PersistentFlags().BoolVar(&app.Debug, DebugName, app.Debug, debugModeUsage)
}

func LoadAppConfig(cmd *cobra.Command) (*config.App, error) {
	app := &config.App{}
	debug, err := cmd.PersistentFlags().GetBool(DebugName)
	if err != nil {
		return nil, err
	}
	app.Debug = debug
	home, err := cmd.PersistentFlags().GetString(HomeDirectoryName)
	if err != nil {
		return nil, err
	}
	app.Home = home
	configFile, err := cmd.PersistentFlags().GetString(ConfigName)
	if err != nil {
		return nil, err
	}
	app.Config = configFile
	return app, nil
}
