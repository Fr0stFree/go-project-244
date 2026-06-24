package cliapp

import (
	"code/internal/parser"
	"errors"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleCLIAppError(t *testing.T) {
	type testCase struct {
		name            string
		err             error
		expectedMessage string
	}

	testCases := []testCase{
		{
			name: "should handle unsupported file type error",
			err: &parser.ParseError{
				Path: "file.txt",
				Err:  parser.ErrUnsupportedFileType,
			},
			expectedMessage: `unable to parse file "file.txt": unsupported file extension`,
		},
		{
			name: "should handle no file extension error",
			err: &parser.ParseError{
				Path: "file",
				Err:  parser.ErrNoFileExtension,
			},
			expectedMessage: `unable to parse file "file": file has no extension`,
		},
		{
			name: "should handle missing file error",
			err: &fs.PathError{
				Op:   "open",
				Path: "file.json",
				Err:  fs.ErrNotExist,
			},
			expectedMessage: `unable to read file "file.json": file does not exist`,
		},
		{
			name: "should handle permission error",
			err: &fs.PathError{
				Op:   "open",
				Path: "file.json",
				Err:  fs.ErrPermission,
			},
			expectedMessage: `unable to read file "file.json": permission denied`,
		},
		{
			name: "should handle unknown parse error",
			err: &parser.ParseError{
				Path: "file.json",
				Err:  errors.New("invalid json"),
			},
			expectedMessage: `unable to parse file "file.json": invalid json`,
		},
		{
			name: "should handle unknown path error",
			err: &fs.PathError{
				Op:   "open",
				Path: "file.json",
				Err:  errors.New("disk error"),
			},
			expectedMessage: `unable to read file "file.json": disk error`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := toCLIExitError(testCase.err)

			require.Error(t, err)
			require.Equal(t, testCase.expectedMessage, err.Error())
		})
	}
}

func TestToCLIExitErrorReturnsUnknownError(t *testing.T) {
	err := errors.New("unexpected error")
	actual := toCLIExitError(err)
	assert.Same(t, err, actual)
}
