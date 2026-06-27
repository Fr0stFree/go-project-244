package cliapp

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v3"
)

func TestRunValidatesPositionalArgumentsCount(t *testing.T) {
	type testCase struct {
		name string
		args []string
	}

	testCases := []testCase{
		{
			name: "should fail without file arguments",
			args: []string{"gendiff"},
		},
		{
			name: "should fail with one file argument",
			args: []string{"gendiff", "file1.json"},
		},
		{
			name: "should fail with extra file argument",
			args: []string{"gendiff", "file1.json", "file2.json", "file3.json"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			cmd := New()
			cmd.ExitErrHandler = func(context.Context, *cli.Command, error) {}

			err := cmd.Run(context.Background(), testCase.args)

			require.Error(t, err)
		})
	}
}
