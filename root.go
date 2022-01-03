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
		//PersistentPreRunE: app.persistentPreRun,
	}

	flags.SetAppFlags(rootCommand)

	app.root = rootCommand

	//loadCommands(
	//	app.b2bCommand,
	//	app.configCommand,
	//	app.pushCommand,
	//	app.reverseCommand,
	//	app.sessionCommand,
	//	app.statusCommand,
	//	app.callbacksCommand,
	//	app.encryptCommand,
	//	app.docsCommand,
	//	app.disburseCommand,
	//)

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
