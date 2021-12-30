package pesakit

import (
	"fmt"
	"github.com/spf13/cobra"
)

func (app *App) reverseCommand() {
	// reverseCmd represents the reverse command
	var reverseCmd = &cobra.Command{
		Use:   "reverse",
		Short: "Reverse the transaction",
		Long:  `Reverse the transaction. Mpesa and Airtel are the only supported networks.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("reverse called")
		},
	}
	app.root.AddCommand(reverseCmd)
}
