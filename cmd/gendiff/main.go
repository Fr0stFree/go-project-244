// Package main is the entry point of the gendiff application.
package main

import (
	"code/internal/cliapp"
	"context"
	"fmt"
	"os"
)

func main() {
	app := cliapp.New()
	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
