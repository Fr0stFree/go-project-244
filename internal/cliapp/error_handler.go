package cliapp

import (
	"code/internal/parser"
	"errors"
	"fmt"
	"io/fs"

	"github.com/urfave/cli/v3"
)

func handleCLIAppError(err error) error {
	var parseErr *parser.ParseError
	if errors.As(err, &parseErr) {
		switch {
		case errors.Is(err, parser.ErrNoFileExtension):
			return cli.Exit(fmt.Sprintf("unable to parse file %q: file has no extension", parseErr.Path), 1)
		case errors.Is(err, parser.ErrUnsupportedFileType):
			return cli.Exit(fmt.Sprintf("unable to parse file %q: unsupported file extension", parseErr.Path), 1)
		default:
			return cli.Exit(parseErr.Error(), 1)
		}
	}

	var pathErr *fs.PathError
	if !errors.As(err, &pathErr) {
		return err
	}

	switch {
	case errors.Is(err, fs.ErrNotExist):
		return cli.Exit(fmt.Sprintf("unable to read file %q: file does not exist", pathErr.Path), 1)
	case errors.Is(err, fs.ErrPermission):
		return cli.Exit(fmt.Sprintf("unable to read file %q: permission denied", pathErr.Path), 1)
	default:
		return cli.Exit(fmt.Sprintf("unable to read file %q: %v", pathErr.Path, pathErr.Err), 1)
	}
}
