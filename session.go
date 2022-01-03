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
