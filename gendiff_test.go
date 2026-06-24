package code

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenDiff(t *testing.T) {
	type testCase struct {
		name        string
		firstPath   string
		secondPath  string
		format      string
		expected    string
		expectedErr string
	}

	testCases := []testCase{
		{
			name:       "should return stylish diff for json files",
			firstPath:  filepath.Join("internal", "testdata", "fixtures", "json", "file1.json"),
			secondPath: filepath.Join("internal", "testdata", "fixtures", "json", "file2.json"),
			format:     "stylish",
			expected:   "{\n  - follow: false\n    host: hexlet.io\n  - proxy: 123.234.53.22\n  - timeout: 50\n  + timeout: 20\n  + verbose: true\n}",
		},
		{
			name:       "should return plain diff for yaml files",
			firstPath:  filepath.Join("internal", "testdata", "fixtures", "yaml", "file1.yaml"),
			secondPath: filepath.Join("internal", "testdata", "fixtures", "yaml", "file2.yaml"),
			format:     "plain",
			expected:   "Property 'follow' was removed\nProperty 'proxy' was removed\nProperty 'timeout' was updated. From 50 to 20\nProperty 'verbose' was added with value: true",
		},
		{
			name:        "should return error for unsupported format",
			format:      "xml",
			expectedErr: "unable to parse style: unsupported format type xml",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := GenDiff(testCase.firstPath, testCase.secondPath, testCase.format)
			if testCase.expectedErr != "" {
				require.Error(t, err)
				assert.Equal(t, testCase.expectedErr, err.Error())

				return
			}

			require.NoError(t, err)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
