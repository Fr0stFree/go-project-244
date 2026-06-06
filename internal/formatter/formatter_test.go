package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	type testCase struct {
		name              string
		outputFormat      string
		expectedFormatter any
		shouldExpectError bool
	}

	testCases := []testCase{
		{
			name:              "should initialize stylish formatter",
			outputFormat:      "stylish",
			expectedFormatter: &stylishDiffFormatter{},
			shouldExpectError: false,
		},
		{
			name:              "should initialize plain formatter",
			outputFormat:      "plain",
			expectedFormatter: &plainDiffFormatter{},
			shouldExpectError: false,
		},
		{
			name:              "should initialize json formatter",
			outputFormat:      "json",
			expectedFormatter: &jsonDiffFormatter{},
			shouldExpectError: false,
		},
		{
			name:              "should return error for unsupported format",
			outputFormat:      "unsupported",
			expectedFormatter: nil,
			shouldExpectError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			formatter, err := New(testCase.outputFormat)

			if testCase.shouldExpectError {
				assert.Error(t, err)
				assert.Nil(t, formatter)
				return
			}

			require.NoError(t, err)
			assert.IsType(t, testCase.expectedFormatter, formatter)
		})
	}
}
