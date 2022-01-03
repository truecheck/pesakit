package pesakit

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (app *App) disburseCommand() {
	// disburseCmd represents the disburse command
	var disburseCmd = &cobra.Command{
		Use:   "disburse",
		Short: "Send money to a mobile money wallet",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("disburse called")
		},
	}
	app.root.AddCommand(disburseCmd)
}
