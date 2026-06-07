package parser

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func fixturePath(paths ...string) string {
	return filepath.Join("..", "..", "testdata", "fixtures", filepath.Join(paths...))
}

func TestParseFile(t *testing.T) {
	type testCase struct {
		name        string
		filepath    string
		expected    map[string]any
		expectedErr error
	}

	testCases := []testCase{
		{
			name:     "should parse JSON file",
			filepath: fixturePath("json", "file1.json"),
			expected: map[string]any{
				"host":    "hexlet.io",
				"timeout": float64(50),
				"proxy":   "123.234.53.22",
				"follow":  false,
			},
		},
		{
			name:     "should parse YAML file",
			filepath: fixturePath("yaml", "file1.yaml"),
			expected: map[string]any{
				"host":    "hexlet.io",
				"timeout": 50,
				"proxy":   "123.234.53.22",
				"follow":  false,
			},
		},
		{
			name:        "should fail on unsupported file type",
			filepath:    fixturePath("toml", "file1.toml"),
			expectedErr: ErrUnsupportedFileType,
		},
		{
			name:        "should fail on missing file extension",
			filepath:    fixturePath("undefined", "file1"),
			expectedErr: ErrNoFileExtension,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			parsed, err := ParseFile(testCase.filepath)
			if testCase.expectedErr != nil {
				require.ErrorIs(t, err, testCase.expectedErr)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, testCase.expected, parsed)
		})
	}
}
