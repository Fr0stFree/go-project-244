package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	type testCase struct {
		name              string
		outputFormat      Style
		expectedFormatter any
	}

	testCases := []testCase{
		{
			name:              "should initialize stylish formatter",
			outputFormat:      Stylish,
			expectedFormatter: &stylishDiffFormatter{},
		},
		{
			name:              "should initialize plain formatter",
			outputFormat:      Plain,
			expectedFormatter: &plainDiffFormatter{},
		},
		{
			name:              "should initialize json formatter",
			outputFormat:      JSON,
			expectedFormatter: &jsonDiffFormatter{},
		},
		{
			name:              "should return stylish formatter for unsupported format",
			outputFormat:      "unsupported",
			expectedFormatter: &stylishDiffFormatter{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			formatter := New(testCase.outputFormat)
			require.IsType(t, testCase.expectedFormatter, formatter)
		})
	}
}

func TestNewStyleFromString(t *testing.T) {
	type testCase struct {
		name          string
		input         string
		expectedStyle Style
		expectedErr   string
	}

	testCases := []testCase{
		{
			name:          "should parse stylish format",
			input:         "stylish",
			expectedStyle: Stylish,
		},
		{
			name:          "should parse plain format",
			input:         "plain",
			expectedStyle: Plain,
		},
		{
			name:          "should parse json format",
			input:         "json",
			expectedStyle: JSON,
		},
		{
			name:        "should fail on unsupported format",
			input:       "xml",
			expectedErr: "unable to parse style: unsupported format type xml",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			style, err := NewStyleFromString(testCase.input)
			if testCase.expectedErr != "" {
				require.Error(t, err)
				assert.Equal(t, testCase.expectedErr, err.Error())

				return
			}

			require.NoError(t, err)
			assert.Equal(t, testCase.expectedStyle, style)
		})
	}
}
