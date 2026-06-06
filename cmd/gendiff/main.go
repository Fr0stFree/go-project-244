// Package main implements the gendiff CLI application.
package main

import (
	"code"
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "file1",
				UsageText: "path to a first file",
			},
			&cli.StringArg{
				Name:      "file2",
				UsageText: "path to a second file",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "format",
				Aliases:  []string{"f"},
				Value:    "stylish",
				Required: false,
				Usage:    "output format",
			},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			firstFile := cmd.StringArg("file1")

			secondFile := cmd.StringArg("file2")
			if firstFile == "" || secondFile == "" {
				return cli.Exit("expected 2 arguments: file1 and file2", 1)
			}

			outputFormat := cmd.String("format")

			result, err := code.GenDiff(firstFile, secondFile, outputFormat)
			if err != nil {
				return err
			}

			fmt.Println(result)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(cmd.ErrWriter, err)
		os.Exit(1)
	}
}
