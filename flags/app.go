package flags

import (
	"github.com/pesakit/pesakit/config"
	"github.com/spf13/cobra"
)

const (
	flagDebugName      = "debug"
	usageDebug         = "debug mode"
	HomeDirectoryName  = "home"
	homeDirectoryUsage = "application home directory"
	ConfigName         = "config"
	configFileUsage    = "config file path"
)

func SetAppFlags(cmd *cobra.Command) {
	app := config.DefaultAppConf()
	cmd.PersistentFlags().StringVarP(&app.Home, HomeDirectoryName, "H", app.Home, homeDirectoryUsage)
	cmd.PersistentFlags().StringVarP(&app.Config, ConfigName, "C", app.Config, configFileUsage)
	cmd.PersistentFlags().BoolVarP(&app.Debug, flagDebugName, "D", app.Debug, usageDebug)
}

func LoadAppConfig(cmd *cobra.Command) (*config.App, error) {
	app := &config.App{}
	cmd = getParentCommand(cmd)
	debug, err := cmd.PersistentFlags().GetBool(flagDebugName)
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

// getParentCommand returns the parent command of the application.
// it takes cmd *cobra.Command as an argument and traverse the tree
// to find the parent command.
func getParentCommand(cmd *cobra.Command) *cobra.Command {
	if cmd.HasParent() {
		return getParentCommand(cmd.Parent())
	}

	return cmd
}
