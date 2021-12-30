package pesakit

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"io"
)

func (app *App) docsCommand() {
	// docsCmd represents the docs command
	var docsCmd = &cobra.Command{
		Use:   "docs",
		Short: "docs generates documentation for pesakit",
		Long: `docs generates documentation for pesakit and saves the documentation
to the specified directory. If no directory is specified the markdown
files will be saved in a /docs path  relative to the current working
directory.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			dir, err := cmd.Flags().GetString("dir")
			if err != nil {
				return err
			}

			// get current working directory and append docs directory
			if dir == "" {
				dir = "docs"
			}
			logger := app.getLogger()
			return docsAction(app.root, logger, dir)
		},
	}

	docsCmd.Flags().StringP("dir", "d", "", "Directory to write the docs to")
	docsCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		markHiddenExcept(app.root.PersistentFlags(), flagDebugMode, flagConfigFile)
		command.Parent().HelpFunc()(command, strings)
	})

	app.root.AddCommand(docsCmd)
}

func docsAction(parentCommand *cobra.Command, out io.Writer, dir string) error {
	if err := doc.GenMarkdownTree(parentCommand, dir); err != nil {
		return err
	}
	_, err := fmt.Fprintf(out, "documentation successfully created in %s\n", dir)
	return err
}
