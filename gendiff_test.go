package code

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func readFileWithExpectedResult(t *testing.T, filename string) string {
	t.Helper()

	data, err := os.ReadFile(filepath.Join("internal", "testdata", "expected", filename))
	require.NoError(t, err)

	return strings.TrimSuffix(string(data), "\n")
}

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
			expected:   readFileWithExpectedResult(t, "file1_file2_stylish.txt"),
		},
		{
			name:       "should return json diff for json files",
			firstPath:  filepath.Join("internal", "testdata", "fixtures", "json", "file1.json"),
			secondPath: filepath.Join("internal", "testdata", "fixtures", "json", "file2.json"),
			format:     "json",
			expected:   readFileWithExpectedResult(t, "file1_file2_json.txt"),
		},
		{
			name:       "should return plain diff for json files",
			firstPath:  filepath.Join("internal", "testdata", "fixtures", "json", "file1.json"),
			secondPath: filepath.Join("internal", "testdata", "fixtures", "json", "file2.json"),
			format:     "plain",
			expected:   readFileWithExpectedResult(t, "file1_file2_plain.txt"),
		},
		{
			name:       "should return stylish diff for yaml files",
			firstPath:  filepath.Join("internal", "testdata", "fixtures", "yaml", "file1.yaml"),
			secondPath: filepath.Join("internal", "testdata", "fixtures", "yaml", "file2.yaml"),
			format:     "stylish",
			expected:   readFileWithExpectedResult(t, "file1_file2_stylish.txt"),
		},
		{
			name:       "should return json diff for yaml files",
			firstPath:  filepath.Join("internal", "testdata", "fixtures", "yaml", "file1.yaml"),
			secondPath: filepath.Join("internal", "testdata", "fixtures", "yaml", "file2.yaml"),
			format:     "json",
			expected:   readFileWithExpectedResult(t, "file1_file2_json.txt"),
		},
		{
			name:       "should return plain diff for yaml files",
			firstPath:  filepath.Join("internal", "testdata", "fixtures", "yaml", "file1.yaml"),
			secondPath: filepath.Join("internal", "testdata", "fixtures", "yaml", "file2.yaml"),
			format:     "plain",
			expected:   readFileWithExpectedResult(t, "file1_file2_plain.txt"),
		},
		{
			name:       "should return stylish diff for complex json files",
			firstPath:  filepath.Join("internal", "testdata", "fixtures", "json", "file3.json"),
			secondPath: filepath.Join("internal", "testdata", "fixtures", "json", "file4.json"),
			format:     "stylish",
			expected:   readFileWithExpectedResult(t, "file3_file4_stylish.txt"),
		},
		{
			name:       "should return stylish diff for complex yaml files",
			firstPath:  filepath.Join("internal", "testdata", "fixtures", "yaml", "file3.yaml"),
			secondPath: filepath.Join("internal", "testdata", "fixtures", "yaml", "file4.yaml"),
			format:     "stylish",
			expected:   readFileWithExpectedResult(t, "file3_file4_stylish.txt"),
		},
		{
			name:       "should return json diff for complex json files",
			firstPath:  filepath.Join("internal", "testdata", "fixtures", "json", "file3.json"),
			secondPath: filepath.Join("internal", "testdata", "fixtures", "json", "file4.json"),
			format:     "json",
			expected:   readFileWithExpectedResult(t, "file3_file4_json.txt"),
		},
		{
			name:       "should return json diff for complex yaml files",
			firstPath:  filepath.Join("internal", "testdata", "fixtures", "yaml", "file3.yaml"),
			secondPath: filepath.Join("internal", "testdata", "fixtures", "yaml", "file4.yaml"),
			format:     "json",
			expected:   readFileWithExpectedResult(t, "file3_file4_json.txt"),
		},
		{
			name:       "should return plain diff for complex json files",
			firstPath:  filepath.Join("internal", "testdata", "fixtures", "json", "file3.json"),
			secondPath: filepath.Join("internal", "testdata", "fixtures", "json", "file4.json"),
			format:     "plain",
			expected:   readFileWithExpectedResult(t, "file3_file4_plain.txt"),
		},
		{
			name:       "should return plain diff for complex yaml files",
			firstPath:  filepath.Join("internal", "testdata", "fixtures", "yaml", "file3.yaml"),
			secondPath: filepath.Join("internal", "testdata", "fixtures", "yaml", "file4.yaml"),
			format:     "plain",
			expected:   readFileWithExpectedResult(t, "file3_file4_plain.txt"),
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
