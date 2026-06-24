// Package cliapp provides the implementation of the CLI application runner for gendiff.
package cliapp

import (
	"code"
	"context"
	"fmt"

	"code/internal/formatter"

	"github.com/urfave/cli/v3"
)

const (
	firstFileArg     = "first-file"
	secondFileArg    = "second-file"
	outputFormatFlag = "format"
)

var cliArgs = []cli.Argument{
	&cli.StringArg{
		Name:      firstFileArg,
		UsageText: "file path to the first configuration file",
	},
	&cli.StringArg{
		Name:      secondFileArg,
		UsageText: "file path to the second configuration file",
	},
}

var cliFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     outputFormatFlag,
		Aliases:  []string{"f"},
		Value:    formatter.Stylish.String(),
		Required: false,
		Usage:    fmt.Sprintf("output format (%s, %s, %s)", formatter.Stylish, formatter.Plain, formatter.JSON),
	},
}

func appAction(_ context.Context, cmd *cli.Command) error {
	firstFile := cmd.StringArg(firstFileArg)
	secondFile := cmd.StringArg(secondFileArg)
	outputFormat := cmd.String(outputFormatFlag)

	result, err := code.GenDiff(firstFile, secondFile, outputFormat)
	if err != nil {
		return toCLIExitError(err)
	}

	fmt.Println(result)

	return nil
}

// New creates and returns a new CLI application.
func New() *cli.Command {
	return &cli.Command{
		Name:      "gendiff",
		Usage:     "Compares two configuration files and shows a difference.",
		Arguments: cliArgs,
		Flags:     cliFlags,
		Action:    appAction,
		UsageText: fmt.Sprintf("gendiff [options] <%s> <%s>", firstFileArg, secondFileArg),
	}
}
