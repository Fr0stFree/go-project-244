package cliapp

import (
	"code/internal/parser"
	"errors"
	"fmt"
	"io/fs"

	"github.com/urfave/cli/v3"
)

const defaultErrExitCode = 1

func toCLIExitError(err error) error {
	var parseErr *parser.ParseError
	if errors.As(err, &parseErr) {
		switch {
		case errors.Is(err, parser.ErrNoFileExtension):
			return cli.Exit(fmt.Sprintf("unable to parse file %q: file has no extension", parseErr.Path), defaultErrExitCode)
		case errors.Is(err, parser.ErrUnsupportedFileType):
			return cli.Exit(fmt.Sprintf("unable to parse file %q: unsupported file extension", parseErr.Path), defaultErrExitCode)
		default:
			return cli.Exit(parseErr.Error(), defaultErrExitCode)
		}
	}

	var pathErr *fs.PathError
	if !errors.As(err, &pathErr) {
		return err
	}

	switch {
	case errors.Is(err, fs.ErrNotExist):
		return cli.Exit(fmt.Sprintf("unable to read file %q: file does not exist", pathErr.Path), defaultErrExitCode)
	case errors.Is(err, fs.ErrPermission):
		return cli.Exit(fmt.Sprintf("unable to read file %q: permission denied", pathErr.Path), defaultErrExitCode)
	default:
		return cli.Exit(fmt.Sprintf("unable to read file %q: %v", pathErr.Path, pathErr.Err), defaultErrExitCode)
	}
}
