package formatter

import (
	"code/internal/diff"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonRender(t *testing.T) {
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
			expected: "{\n  \"foo\": {\n    \"type\": \"added\",\n    \"value\": \"bar\"\n  }\n}",
		},
		{
			name: "should render removed record",
			records: []diff.Record{
				{Key: "foo", State: diff.Removed, OldValue: "bar", NewValue: nil, Children: nil},
			},
			expected: "{\n  \"foo\": {\n    \"type\": \"removed\",\n    \"value\": \"bar\"\n  }\n}",
		},
		{
			name: "should render changed record",
			records: []diff.Record{
				{Key: "foo", State: diff.Changed, OldValue: "bar", NewValue: "baz", Children: nil},
			},
			expected: "{\n  \"foo\": {\n    \"new\": \"baz\",\n    \"old\": \"bar\",\n    \"type\": \"changed\"\n  }\n}",
		},
		{
			name: "should render unchanged record",
			records: []diff.Record{
				{Key: "foo", State: diff.Unchanged, OldValue: "bar", NewValue: "bar", Children: nil},
			},
			expected: "{\n  \"foo\": {\n    \"type\": \"unchanged\",\n    \"value\": \"bar\"\n  }\n}",
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
			expected: "{\n  \"foo\": {\n    \"bar\": {\n      \"type\": \"added\",\n      \"value\": \"baz\"\n    }\n  }\n}",
		},
		{
			name: "should render unchanged records inside nested records",
			records: []diff.Record{
				{
					Key: "common", State: diff.Nested, OldValue: nil, NewValue: nil,
					Children: []diff.Record{
						{
							Key: "setting6", State: diff.Nested, OldValue: nil, NewValue: nil,
							Children: []diff.Record{
								{Key: "key", State: diff.Unchanged, OldValue: "value", NewValue: "value", Children: nil},
								{Key: "ops", State: diff.Added, OldValue: nil, NewValue: "vops", Children: nil},
							},
						},
					},
				},
			},
			expected: "{\n  \"common\": {\n    \"setting6\": {\n      \"key\": {\n        \"type\": \"unchanged\",\n        \"value\": \"value\"\n      },\n      \"ops\": {\n        \"type\": \"added\",\n        \"value\": \"vops\"\n      }\n    }\n  }\n}",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			formatter := &jsonDiffFormatter{}
			result, err := formatter.Render(testCase.records)
			assert.NoError(t, err)
			assert.Equal(t, testCase.expected, result)
		})
	}
}
