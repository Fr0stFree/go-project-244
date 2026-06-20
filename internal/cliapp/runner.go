// Package cliapp provides the implementation of the CLI application runner for gendiff.
package cliapp

import (
	"context"
	"fmt"

	"code/internal/formatter"

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
		UsageText: "gendiff [options] <file1> <file2>",
	}
}

func makeFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "format",
			Aliases:  []string{"f"},
			Value:    formatter.Stylish.String(),
			Required: false,
			Usage:    fmt.Sprintf("output format (%s, %s, %s)", formatter.Stylish, formatter.Plain, formatter.JSON),
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
