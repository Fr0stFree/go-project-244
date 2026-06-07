package parser

import "errors"

var (
	// ErrUnsupportedFileType means the file extension is not supported by parser.
	ErrUnsupportedFileType = errors.New("unable to parse file: unsupported extension")
	// ErrNoFileExtension means the file path has no extension to select parser from.
	ErrNoFileExtension = errors.New("unable to parse file: file has no extension")
)

// ParseError adds file path context to lower-level parse errors.
type ParseError struct {
	Err  error
	Path string
}

func (e ParseError) Error() string {
	return e.Err.Error()
}

func (e ParseError) Unwrap() error {
	return e.Err
}

func newParseError(err error, path string) *ParseError {
	return &ParseError{Err: err, Path: path}
}
