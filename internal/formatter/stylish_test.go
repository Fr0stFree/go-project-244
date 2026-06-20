package formatter

import (
	"code/internal/diff"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStylishRender(t *testing.T) {
	type testCase struct {
		name     string
		records  []diff.Record
		expected string
	}

	testCases := []testCase{
		{
			name: "should render added record",
			records: []diff.Record{
				{Key: "foo", State: diff.Added, OldValue: nil, NewValue: "bar", Children: nil},
			},
			expected: "{\n  + foo: bar\n}",
		},
		{
			name: "should render removed record",
			records: []diff.Record{
				{Key: "foo", State: diff.Removed, OldValue: "bar", NewValue: nil, Children: nil},
			},
			expected: "{\n  - foo: bar\n}",
		},
		{
			name: "should render changed record",
			records: []diff.Record{
				{Key: "foo", State: diff.Changed, OldValue: "bar", NewValue: "baz", Children: nil},
			},
			expected: "{\n  - foo: bar\n  + foo: baz\n}",
		},
		{
			name: "should render unchanged record",
			records: []diff.Record{
				{Key: "foo", State: diff.Unchanged, OldValue: "bar", NewValue: "bar", Children: nil},
			},
			expected: "{\n    foo: bar\n}",
		},
		{
			name: "should render nested records",
			records: []diff.Record{
				{
					Key: "foo", State: diff.Nested, OldValue: nil, NewValue: nil,
					Children: []diff.Record{
						{Key: "bar", State: diff.Added, OldValue: nil, NewValue: "baz", Children: nil},
					},
				},
			},
			expected: "{\n    foo: {\n      + bar: baz\n    }\n}",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			formatter := &stylishDiffFormatter{}
			result, err := formatter.Render(testCase.records)
			assert.NoError(t, err)
			assert.Equal(t, testCase.expected, result)
		})
	}
}
