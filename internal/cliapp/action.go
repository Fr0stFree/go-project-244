package cliapp

import (
	"code"
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func cliAppAction(onError func(error) error) func(_ context.Context, cmd *cli.Command) error {
	return func(_ context.Context, cmd *cli.Command) error {
		files := cmd.StringArgs("files")
		outputFormat := cmd.String("format")

		result, err := code.GenDiff(files[0], files[1], outputFormat)
		if err != nil {
			return onError(err)
		}

		fmt.Println(result)

		return nil
	}
}
