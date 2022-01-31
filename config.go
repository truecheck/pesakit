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

package main

import (
	"fmt"
	"github.com/pesakit/pesakit/flags"
	"github.com/spf13/cobra"
)

func (app *App) configCommand() {
	rootCommand := app.root
	configCommand := &cobra.Command{
		Use:   "config",
		Short: "manages clients configuration",
		Long: `manages clients configuration. This command is used to create, view,update and delete
clients configurations`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	setFlagsFunc := func() {
		flags.SetMno(configCommand, flags.PERSISTENT)
		flags.SetMpesa(configCommand)
		flags.SetTigoPesa(configCommand)
	}
	setFlagsFunc()
	app.configPrintCommand(configCommand)
	rootCommand.AddCommand(configCommand)
}

func (app *App) configPrintCommand(parent *cobra.Command) {
	configPrintCommand := &cobra.Command{
		Use:   "print",
		Short: "prints the configuration",
		Long:  `prints the configuration if not any is specified all configurations are printed`,
		Run: func(cmd *cobra.Command, args []string) {
			writer := app.getWriter()
			fmt.Fprintf(writer, "Printing configurations\n")
			fmt.Fprintf(writer, "======================\n")
			fmt.Fprintf(writer, "mpesa: %+v\n", app.mpesa.Config)
			fmt.Fprintf(writer, "tigo-pesa: %+v\n", app.tigo.Config)
		},
	}

	//configPrintCommand.SetHelpFunc(func(command *cobra.Command, strings []string) {
	//	set := parentCommand.Parent().PersistentFlags()
	//	markHiddenExcept(set, flagDebugMode, flagConfigAirtel, flagConfigMpesa, flagConfigTigo)
	//	command.Parent().HelpFunc()(command, strings)
	//})

	parent.AddCommand(configPrintCommand)
}
