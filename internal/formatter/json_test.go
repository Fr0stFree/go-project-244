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
			expected: "[\n  {\n    \"key\": \"foo\",\n    \"type\": \"added\",\n    \"new\": \"bar\"\n  }\n]",
		},
		{
			name: "should render removed record",
			records: []diff.Record{
				{Key: "foo", State: diff.Removed, OldValue: "bar", NewValue: nil, Children: nil},
			},
			expected: "[\n  {\n    \"key\": \"foo\",\n    \"type\": \"removed\",\n    \"old\": \"bar\"\n  }\n]",
		},
		{
			name: "should render changed record",
			records: []diff.Record{
				{Key: "foo", State: diff.Changed, OldValue: "bar", NewValue: "baz", Children: nil},
			},
			expected: "[\n  {\n    \"key\": \"foo\",\n    \"type\": \"changed\",\n    \"old\": \"bar\",\n    \"new\": \"baz\"\n  }\n]",
		},
		{
			name: "should render unchanged record",
			records: []diff.Record{
				{Key: "foo", State: diff.Unchanged, OldValue: "bar", NewValue: "bar", Children: nil},
			},
			expected: "[\n  {\n    \"key\": \"foo\",\n    \"type\": \"unchanged\",\n    \"old\": \"bar\",\n    \"new\": \"bar\"\n  }\n]",
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
			expected: "[\n  {\n    \"key\": \"foo\",\n    \"type\": \"nested\",\n    \"children\": [\n      {\n        \"key\": \"bar\",\n        \"type\": \"added\",\n        \"new\": \"baz\"\n      }\n    ]\n  }\n]",
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
			expected: "[\n  {\n    \"key\": \"common\",\n    \"type\": \"nested\",\n    \"children\": [\n      {\n        \"key\": \"setting6\",\n        \"type\": \"nested\",\n        \"children\": [\n          {\n            \"key\": \"key\",\n            \"type\": \"unchanged\",\n            \"old\": \"value\",\n            \"new\": \"value\"\n          },\n          {\n            \"key\": \"ops\",\n            \"type\": \"added\",\n            \"new\": \"vops\"\n          }\n        ]\n      }\n    ]\n  }\n]",
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
