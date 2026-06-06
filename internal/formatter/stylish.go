package formatter

import (
	"code/internal/diff"
	"fmt"
	"slices"
	"strings"

	"github.com/samber/lo"
)

const indentSize int = 4
const markerSize int = 2

type stylishDiffFormatter struct{}

// Render converts a diff into the stylish string representation.
func (s *stylishDiffFormatter) Render(records []diff.Record) string {
	return s.renderRecords(records, 0)
}

func (s *stylishDiffFormatter) renderRecords(records []diff.Record, depth int) string {
	builder := strings.Builder{}

	currentIndent := strings.Repeat(" ", depth*indentSize)
	nextIndent := strings.Repeat(" ", (depth+1)*indentSize)
	markerIndent := strings.Repeat(" ", (depth+1)*indentSize-markerSize)

	builder.WriteRune('{')
	builder.WriteRune('\n')

	for _, record := range records {
		switch record.State {
		case diff.Added:
			fmt.Fprintf(&builder, "%s%c %s: %s\n", markerIndent, '+', record.Key, s.formatValue(record.NewValue, depth+1))
		case diff.Removed:
			fmt.Fprintf(&builder, "%s%c %s: %s\n", markerIndent, '-', record.Key, s.formatValue(record.OldValue, depth+1))
		case diff.Unchanged:
			fmt.Fprintf(&builder, "%s%c %s: %s\n", markerIndent, ' ', record.Key, s.formatValue(record.OldValue, depth+1))
		case diff.Changed:
			fmt.Fprintf(&builder, "%s%c %s: %s\n", markerIndent, '-', record.Key, s.formatValue(record.OldValue, depth+1))
			fmt.Fprintf(&builder, "%s%c %s: %s\n", markerIndent, '+', record.Key, s.formatValue(record.NewValue, depth+1))
		case diff.Nested:
			fmt.Fprintf(&builder, "%s%s: %s\n", nextIndent, record.Key, s.renderRecords(record.Children, depth+1))
		default:
			panic(fmt.Sprintf("unknown diff record state: %v", record.State))
		}
	}

	builder.WriteString(currentIndent)
	builder.WriteRune('}')

	return builder.String()
}

func (s *stylishDiffFormatter) formatValue(unknownValue any, depth int) string {
	switch value := unknownValue.(type) {
	case map[string]any:
		builder := strings.Builder{}

		currentIndent := strings.Repeat(" ", depth*indentSize)
		nextIndent := strings.Repeat(" ", (depth+1)*indentSize)

		builder.WriteRune('{')
		builder.WriteRune('\n')

		keys := lo.Keys(value)
		slices.Sort(keys)

		for _, key := range keys {
			v := value[key]
			fmt.Fprintf(&builder, "%s%s: %s\n", nextIndent, key, s.formatValue(v, depth+1))
		}

		builder.WriteString(currentIndent)
		builder.WriteRune('}')

		return builder.String()
	case nil:
		return "null"
	default:
		return fmt.Sprintf("%v", value)
	}
}
