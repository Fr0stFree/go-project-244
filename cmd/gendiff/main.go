package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
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
			args := cmd.Args()
			if args.Len() != 2 {
				return fmt.Errorf("expected 2 arguments: file1 and file2")
			}
			firstFP := args.Get(0)
			secondFP := args.Get(1)
			format := cmd.String("format")
			fmt.Printf(
				"file1=%s file2=%s format=%s\n",
				firstFP,
				secondFP,
				format,
			)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
