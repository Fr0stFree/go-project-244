// Package cliapp provides the implementation of the CLI application runner for gendiff.
package cliapp

import (
	"context"

	"github.com/urfave/cli/v3"
)

type runner interface {
	Run(context.Context, []string) error
}

// NewRunner creates and returns a new CLI application runner.
func NewRunner() runner {
	return &cli.Command{
		Name:      "gendiff",
		Usage:     "Compares two configuration files and shows a difference.",
		Arguments: makeArguments(),
		Flags:     makeFlags(),
		Action:    cliAppAction(handleCLIAppError),
	}
}

func makeFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "format",
			Aliases:  []string{"f"},
			Value:    "stylish",
			Required: false,
			Usage:    "output format",
		},
	}
}

func makeArguments() []cli.Argument {
	return []cli.Argument{
		&cli.StringArgs{
			Name:      "files",
			UsageText: "two files to compare",
			Min:       2,
			Max:       2,
		},
	}

}
