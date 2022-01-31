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
	"github.com/pesakit/pesakit/flags"
	"github.com/pesakit/pesakit/mno"
	"github.com/spf13/cobra"
)

func (app *App) sessionCommand() {
	var sessionCmd = &cobra.Command{
		Use:   "session",
		Short: "generates session id for a particular mno",
		Long: `generates session id for a particular mno. The sessionId is then used
to authenticate the user. For-example, In Mpesa API it is used as
bearer token. So after generating the sessionId, the user can
use it with other API client tools like Postman.

You need to specify the mno if not specified, the default mno is mpesa.
`,
		Run: func(cmd *cobra.Command, args []string) {
			mnoValue, err := flags.GetMno(cmd, flags.PERSISTENT)
			if err != nil {
				return
			}

			switch mnoValue {
			case mno.Airtel:
				app.Logger().Printf("retrieving session id for airtel\n")
				return
			case mno.Vodacom:
				app.Logger().Printf("retrieving session id for vodacom\n")
				return
			case mno.Tigo:
				app.Logger().Printf("retrieving session id for tigo\n")
				return

			default:
				_ = cmd.Help()
			}

		},
	}
	flags.SetMno(sessionCmd, flags.PERSISTENT)
	app.root.AddCommand(sessionCmd)
}
