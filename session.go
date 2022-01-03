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
	"github.com/spf13/cobra"
)

const (
	flagSessionMno  = "mno"
	pFlagSessionMno = "m"
	defSessionValue = "mpesa"
	usageSessionMno = "mobile money operator"
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
			varSessionMno, err := cmd.Flags().GetString(flagSessionMno)
			if err != nil {
				_, _ = fmt.Fprintf(app.getWriter(), "error loading the mno: %v\n", err)
			}
			if varSessionMno == "" {
				varSessionMno = defSessionValue
			} else if varSessionMno == "mpesa" || varSessionMno == "vodacom" {
				_, _ = fmt.Fprintln(app.getWriter(), "mpesa session id")
			} else if varSessionMno == "airtel" {
				_, _ = fmt.Fprintln(app.getWriter(), "airtel session id")
			} else if varSessionMno == "tigo" {
				_, _ = fmt.Fprintln(app.getWriter(), "tigo session id")
			} else {
				_, _ = fmt.Fprintf(app.getWriter(), "unknown mno: %v\n", varSessionMno)
				_ = cmd.Help()

				return

			}

		},
	}
	sessionCmd.Flags().StringP(flagSessionMno, pFlagSessionMno, defSessionValue, usageSessionMno)
	app.root.AddCommand(sessionCmd)
}
