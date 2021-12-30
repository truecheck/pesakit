package pesakit

import (
	"fmt"
	"github.com/spf13/cobra"
)

func (app *App) statusCommand() {
	// statusCmd represents the status command
	var statusCmd = &cobra.Command{
		Use:   "status",
		Short: "check status of transaction",
		Long: `check status of transaction. Only Airtel Money and Mpesa transactions
are supported.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("status called")
		},
	}
	app.root.AddCommand(statusCmd)
}
