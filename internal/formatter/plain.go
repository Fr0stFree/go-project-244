package formatter

import (
	"code/internal/diff"
	"fmt"
	"strings"
)

type plainDiffFormatter struct{}

func (p *plainDiffFormatter) Render(records []diff.Record) string {
	return strings.TrimSuffix(p.renderRecords(records, ""), "\n")
}

func (p *plainDiffFormatter) renderRecords(records []diff.Record, parentKey string) string {
	builder := strings.Builder{}

	for _, record := range records {
		currentKey := record.Key
		if parentKey != "" {
			currentKey = parentKey + "." + record.Key
		}

		switch record.State {
		case diff.Added:
			fmt.Fprintf(&builder, "Property '%s' was added with value: %s\n", currentKey, p.formatValue(record.NewValue))
		case diff.Removed:
			fmt.Fprintf(&builder, "Property '%s' was removed\n", currentKey)
		case diff.Changed:
			fmt.Fprintf(&builder, "Property '%s' was updated. From %s to %s\n", currentKey, p.formatValue(record.OldValue), p.formatValue(record.NewValue))
		case diff.Unchanged:
			continue
		case diff.Nested:
			builder.WriteString(p.renderRecords(record.Children, currentKey))
		default:
			panic(fmt.Sprintf("unknown diff record state: %v", record.State))
		}
	}

	return builder.String()
}
func (p *plainDiffFormatter) formatValue(unknownValue any) string {
	switch value := unknownValue.(type) {
	case map[string]any:
		return "[complex value]"
	case string:
		return fmt.Sprintf("'%s'", value)
	case nil:
		return "null"
	default:
		return fmt.Sprintf("%v", value)
	}
}
