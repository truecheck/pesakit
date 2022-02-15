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
	"github.com/urfave/cli/v2"
)

var (
	reverseCommandName  = "reverse"
	reverseCommandUsage = "reverse a transaction"
)

func reverseCommand() *cli.Command {
	return &cli.Command{
		Name:         reverseCommandName,
		Aliases:      nil,
		Usage:        reverseCommandUsage,
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action: func(c *cli.Context) error {
			return nil
		},
		OnUsageError:           nil,
		Subcommands:            nil,
		Flags:                  nil,
		SkipFlagParsing:        false,
		HideHelp:               false,
		HideHelpCommand:        false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}
}

//func (app *App) reverseCommand() {
//	// reverseCmd represents the reverse command
//	var reverseCmd = &cobra.Command{
//		Use:   "reverse",
//		Short: "Reverse the transaction",
//		Long:  `Reverse the transaction. Mpesa and Airtel are the only supported networks.`,
//		Run: func(cmd *cobra.Command, args []string) {
//			fmt.Println("reverse called")
//		},
//	}
//	app.root.AddCommand(reverseCmd)
//}
