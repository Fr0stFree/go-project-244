package cliapp

import (
	"code"
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func cliAppAction(onError func(error) error) func(_ context.Context, cmd *cli.Command) error {
	return func(_ context.Context, cmd *cli.Command) error {
		firstFile := cmd.StringArg("file1")

		secondFile := cmd.StringArg("file2")
		if firstFile == "" || secondFile == "" {
			return cli.Exit("expected 2 arguments: file1 and file2", 1)
		}

		outputFormat := cmd.String("format")

		result, err := code.GenDiff(firstFile, secondFile, outputFormat)
		if err != nil {
			return onError(err)
		}

		fmt.Println(result)

		return nil
	}
}
