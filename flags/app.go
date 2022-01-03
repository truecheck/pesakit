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
