package pesakit

import (
	"fmt"
	"github.com/spf13/cobra"
)

func (app *App) sessionCommand() {
	// sessionCmd represents the session command
	var sessionCmd = &cobra.Command{
		Use:   "session",
		Short: "generates session id for a particular mno",
		Long: `generates session id for a particular mno. The sessionId is then used
to authenticate the user. For-example, In Mpesa API it is used as
bearer token. So after generating the sessionId, the user can
use it with other API client tools like Postman.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("session called")
		},
	}
	app.root.AddCommand(sessionCmd)
}
