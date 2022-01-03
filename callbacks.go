/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package pesakit

import (
	"github.com/pesakit/pesakit/flags"
	"github.com/pesakit/pesakit/mno"
	"github.com/spf13/cobra"
)

const (
	push callbackOperation = "push"
)

type callbackOperation string

type callbackParams struct {
	Mno       mno.Mno
	Host      string
	Port      int
	Path      string
	Operation callbackOperation
}

func loadCallbackParams(cmd *cobra.Command) (*callbackParams, error) {
	host, err := cmd.Flags().GetString("host")
	if err != nil {
		return nil, err
	}
	port, err := cmd.Flags().GetInt("port")
	if err != nil {
		return nil, err
	}
	path, err := cmd.Flags().GetString("path")
	if err != nil {
		return nil, err
	}
	operation, err := cmd.Flags().GetString("operation")
	if err != nil {
		return nil, err
	}
	mnoValue, err := flags.LoadMnoConfig(cmd, flags.PERSISTENT)
	if err != nil {
		return nil, err
	}

	return &callbackParams{
		Mno:       mnoValue,
		Host:      host,
		Port:      port,
		Path:      path,
		Operation: callbackOperation(operation),
	}, nil

}

func (app *App) callbacksCommand() {

	// callbacksCmd represents the callbacks command
	var callbacksCmd = &cobra.Command{
		Use:   "callbacks",
		Short: "Monitor http callbacks from mobile money providers",
		Long:  `Monitor http callbacks from mobile money providers.`,
		Run: func(cmd *cobra.Command, args []string) {
			if app.getDebugMode() {
				app.Logger().Printf("debug mode is ON\n")
				app.Logger().Printf("callbacks called\n")
			}

			params, err := loadCallbackParams(cmd)
			if err != nil {
				app.Logger().Printf("error: %s\n", err)

				return
			}

			app.Logger().Printf("callback params: %+v\n", params)

		},
	}
	flags.SetMnoFlag(callbacksCmd, flags.PERSISTENT)
	callbacksCmd.PersistentFlags().Int("port", 8080, "callback server port")
	callbacksCmd.PersistentFlags().String("host", "localhost", "callback server host")
	callbacksCmd.PersistentFlags().String("path", "/callbacks", "callback server path")
	callbacksCmd.PersistentFlags().String("operation", "push", "operation to listen to")
	app.root.AddCommand(callbacksCmd)
}
