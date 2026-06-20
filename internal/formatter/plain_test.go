package formatter

import (
	"code/internal/diff"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlainRender(t *testing.T) {
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
			expected: "Property 'foo' was added with value: 'bar'",
		},
		{
			name: "should render removed record",
			records: []diff.Record{
				{Key: "foo", State: diff.Removed, OldValue: "bar", NewValue: nil, Children: nil},
			},
			expected: "Property 'foo' was removed",
		},
		{
			name: "should render changed record",
			records: []diff.Record{
				{Key: "foo", State: diff.Changed, OldValue: "bar", NewValue: "baz", Children: nil},
			},
			expected: "Property 'foo' was updated. From 'bar' to 'baz'",
		},
		{
			name: "should render unchanged record",
			records: []diff.Record{
				{Key: "foo", State: diff.Unchanged, OldValue: "bar", NewValue: "bar", Children: nil},
			},
			expected: "",
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
			expected: "Property 'foo.bar' was added with value: 'baz'",
		},
		{
			name: "should render complex value",
			records: []diff.Record{
				{Key: "foo", State: diff.Added, OldValue: nil, NewValue: map[string]any{"bar": "baz"}, Children: nil},
			},
			expected: "Property 'foo' was added with value: [complex value]",
		},
		{
			name: "should render null value",
			records: []diff.Record{
				{Key: "foo", State: diff.Added, OldValue: nil, NewValue: nil, Children: nil},
			},
			expected: "Property 'foo' was added with value: null",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			formatter := &plainDiffFormatter{}
			result, err := formatter.Render(testCase.records)
			assert.NoError(t, err)
			assert.Equal(t, testCase.expected, result)
		})
	}
}
