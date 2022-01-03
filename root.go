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
	"github.com/pesakit/pesakit/flags"
	"github.com/spf13/cobra"
)

type Config struct {
}

func (app *App) createRootCommand() {

	var rootCommand = &cobra.Command{
		Use:   appName,
		Short: appShortDesc,
		Long:  appLongDescription,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				app.Logger().Fatal(err)
			}
		},
		//PersistentPreRunE: app.persistentPreRun,
	}

	flags.SetAppFlags(rootCommand)
	rootCommand.PersistentFlags().String("test", "", "test")
	rootCommand.Flags().String("test", "", "test")
	app.root = rootCommand

	addSubCommandsFunc := func() {
		app.callbacksCommand()
		app.configCommand()
		app.pushCommand()
		app.sessionCommand()

	}

	addSubCommandsFunc()
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
