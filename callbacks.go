/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package pesakit

import (
	"fmt"
	"github.com/spf13/cobra"
)

func (app *App) callbacksCommand() {

	// callbacksCmd represents the callbacks command
	var callbacksCmd = &cobra.Command{
		Use:   "callbacks",
		Short: "Monitor http callbacks from mobile money providers",
		Long:  `Monitor http callbacks from mobile money providers.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("callbacks called")
		},
	}
	markHiddenExcept(app.root.PersistentFlags(), "help")
	app.root.AddCommand(callbacksCmd)
}
