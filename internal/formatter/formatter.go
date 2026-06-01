// Package formatter converts differences into output formats.
package formatter

import (
	"code/internal/diff"
	"fmt"
)

type diffFormatter interface {
	Render([]diff.Node) string
}

// New creates a formatter for the given format type.
func New(fmtType string) (diffFormatter, error) {
	switch fmtType {
	case "stylish":
		return &stylishDiffFormatter{}, nil
	default:
		return nil, fmt.Errorf("unable to format: unsupported format type %s", fmtType)
	}
}
