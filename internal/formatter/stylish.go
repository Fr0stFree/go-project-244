package formatter

import (
	"code/internal/diff"
	"fmt"
	"strings"
)

type stylishDiffFormatter struct{}

// Render converts a diff into the stylish string representation.
func (s *stylishDiffFormatter) Render(nodes []diff.Node) string {
	builder := strings.Builder{}
	builder.WriteRune('{')
	builder.WriteRune('\n')

	for _, node := range nodes {
		switch node.State {
		case diff.Added:
			fmt.Fprintf(&builder, " + %s: %v\n", node.Key, node.NewValue)
		case diff.Removed:
			fmt.Fprintf(&builder, " - %s: %v\n", node.Key, node.OldValue)
		case diff.Unchanged:
			fmt.Fprintf(&builder, "   %s: %v\n", node.Key, node.OldValue)
		case diff.Changed:
			fmt.Fprintf(&builder, " - %s: %v\n", node.Key, node.OldValue)
			fmt.Fprintf(&builder, " + %s: %v\n", node.Key, node.NewValue)
		}
	}

	builder.WriteRune('}')

	return builder.String()
}
