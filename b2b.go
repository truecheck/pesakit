package pesakit

import (
	"fmt"
	"github.com/spf13/cobra"
)

func (app *App) b2bCommand() {

	// b2bCmd represents the b2b command
	var b2bCmd = &cobra.Command{
		Use:   "b2b",
		Short: "sends money from one business account to another",
		Long: `sends money from one business account to another. This is like normal 
disbursement except that it happens between two business accounts`,

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("b2b called")
		},
	}

	app.root.AddCommand(b2bCmd)
}
