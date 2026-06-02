package formatter

import (
	"code/internal/diff"
	"fmt"
	"strings"
)

type stylishDiffFormatter struct{}

// Render converts a diff into the stylish string representation.
func (s *stylishDiffFormatter) Render(records []diff.Record) string {
	builder := strings.Builder{}
	builder.WriteRune('{')
	builder.WriteRune('\n')

	for _, record := range records {
		switch record.State {
		case diff.Added:
			fmt.Fprintf(&builder, " + %s: %v\n", record.Key, record.NewValue)
		case diff.Removed:
			fmt.Fprintf(&builder, " - %s: %v\n", record.Key, record.OldValue)
		case diff.Unchanged:
			fmt.Fprintf(&builder, "   %s: %v\n", record.Key, record.OldValue)
		case diff.Changed:
			fmt.Fprintf(&builder, " - %s: %v\n", record.Key, record.OldValue)
			fmt.Fprintf(&builder, " + %s: %v\n", record.Key, record.NewValue)
		}
	}

	builder.WriteRune('}')

	return builder.String()
}
