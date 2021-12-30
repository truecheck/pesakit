package main

import (
	"fmt"
	"os"

	"github.com/pesakit/pesakit"
	"github.com/pesakit/pesakit/env"
)

func main() {
	env.LoadConfigFrom(".env")
	app := pesakit.New()
	if err := app.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
