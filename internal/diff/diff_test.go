package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	type testCase struct {
		name       string
		inputLeft  map[string]any
		inputRight map[string]any
		expected   []Record
	}

	testCases := []testCase{
		{
			name:       "should identify changed record",
			inputLeft:  map[string]any{"foo": "bar"},
			inputRight: map[string]any{"foo": "baz"},
			expected:   []Record{{"foo", Changed, "bar", "baz", nil}},
		},
		{
			name:       "should identify unchanged record",
			inputLeft:  map[string]any{"foo": "bar"},
			inputRight: map[string]any{"foo": "bar"},
			expected:   []Record{{"foo", Unchanged, "bar", "bar", nil}},
		},
		{
			name:       "should identify added record",
			inputLeft:  map[string]any{},
			inputRight: map[string]any{"foo": "bar"},
			expected:   []Record{{"foo", Added, nil, "bar", nil}},
		},
		{
			name:       "should identify removed record",
			inputLeft:  map[string]any{"foo": "bar"},
			inputRight: map[string]any{},
			expected:   []Record{{"foo", Removed, "bar", nil, nil}},
		},
		{
			name:       "should identify multiple records",
			inputLeft:  map[string]any{"foo": "bar", "one": "two"},
			inputRight: map[string]any{"foo": "baz", "x": "y"},
			expected:   []Record{{"foo", Changed, "bar", "baz", nil}, {"one", Removed, "two", nil, nil}, {"x", Added, nil, "y", nil}},
		},
		{
			name:       "should handle empty inputs",
			inputLeft:  map[string]any{},
			inputRight: map[string]any{},
			expected:   []Record{},
		},
		{
			name:       "should handle nested maps",
			inputLeft:  map[string]any{"foo": map[string]any{"bar": "baz"}},
			inputRight: map[string]any{"foo": map[string]any{"bar": "qux"}},
			expected:   []Record{{"foo", Nested, nil, nil, []Record{{"bar", Changed, "baz", "qux", nil}}}},
		},
		{
			name:       "should build nested record for equal maps",
			inputLeft:  map[string]any{"settings": map[string]any{"timeout": 50}},
			inputRight: map[string]any{"settings": map[string]any{"timeout": 50}},
			expected: []Record{
				{"settings", Nested, nil, nil, []Record{{"timeout", Unchanged, 50, 50, nil}}},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := Build(testCase.inputLeft, testCase.inputRight)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}
