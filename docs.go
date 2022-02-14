/*
 * MIT License
 *
 * Copyright (c) 2021 PESAKIT - MOBILE MONEY TOOLBOX
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

package main

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
			logger := app.getWriter()
			return docsAction(app.root, logger, dir)
		},
	}

	docsCmd.Flags().StringP("dir", "d", "", "Directory to write the docs to")

	app.root.AddCommand(docsCmd)
}

func docsAction(parentCommand *cobra.Command, out io.Writer, dir string) error {
	if err := doc.GenMarkdownTree(parentCommand, dir); err != nil {
		return err
	}
	_, err := fmt.Fprintf(out, "documentation successfully created in %s\n", dir)
	return err
}
