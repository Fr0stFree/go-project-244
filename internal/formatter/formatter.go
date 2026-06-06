// Package formatter converts differences into output formats.
package formatter

import (
	"code/internal/diff"
	"fmt"
)

type diffFormatter interface {
	Render([]diff.Record) string
}

// New creates a formatter for the given format type.
func New(outputFormat string) (diffFormatter, error) {
	switch outputFormat {
	case "stylish":
		return &stylishDiffFormatter{}, nil
	case "plain":
		return &plainDiffFormatter{}, nil
	case "json":
		return &jsonDiffFormatter{}, nil
	default:
		return nil, fmt.Errorf("unable to format: unsupported format type %s", outputFormat)
	}
}
