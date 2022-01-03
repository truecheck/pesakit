package pesakit

import (
	"github.com/pesakit/pesakit/flags"
	"github.com/spf13/cobra"
)

type Config struct {
}

func (app *App) createRootCommand() {

	var rootCommand = &cobra.Command{
		Use:   appName,
		Short: appShortDesc,
		Long:  appLongDescription,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				app.Logger().Fatal(err)
			}
		},
		//PersistentPreRunE: app.persistentPreRun,
	}

	flags.SetAppFlags(rootCommand)
	rootCommand.PersistentFlags().String("test", "", "test")
	rootCommand.Flags().String("test", "", "test")
	app.root = rootCommand

	app.callbacksCommand()
}

func loadCommands(fns ...func()) {
	for _, fn := range fns {
		fn()
	}
}

// getParentCommand returns the parent command of the application.
// it takes cmd *cobra.Command as an argument and traverse the tree
// to find the parent command.
func getParentCommand(cmd *cobra.Command) *cobra.Command {
	if cmd.HasParent() {
		return getParentCommand(cmd.Parent())
	}
	return cmd
}
