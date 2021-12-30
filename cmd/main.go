package main

import (
	"fmt"
	"github.com/pesakit/pesakit"
	"os"
)

func main() {
	app := pesakit.New()
	if err := app.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
