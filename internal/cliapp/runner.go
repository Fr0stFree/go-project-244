// Package cliapp provides the implementation of the CLI application runner for gendiff.
package cliapp

import (
	"code"
	"context"
	"fmt"

	"code/internal/formatter"

	"github.com/urfave/cli/v3"
)

// New creates and returns a new CLI application.
func New() *cli.Command {
	return &cli.Command{
		Name:      "gendiff",
		Usage:     "Compares two configuration files and shows a difference.",
		Arguments: makeArguments(),
		Flags:     makeFlags(),
		Action:    appAction,
		UsageText: "gendiff [options] <file1> <file2>",
	}
}

func appAction(_ context.Context, cmd *cli.Command) error {
	files := cmd.StringArgs("files")
	outputFormat := cmd.String("format")

	result, err := code.GenDiff(files[0], files[1], outputFormat)
	if err != nil {
		return toCLIExitError(err)
	}

	fmt.Println(result)

	return nil
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
