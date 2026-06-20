// Package formatter converts differences into output formats.
package formatter

import (
	"code/internal/diff"
	"fmt"
)

// Style represents the output format type.
type Style string

// Supported output format types.
const (
	Stylish Style = "stylish"
	Plain   Style = "plain"
	JSON    Style = "json"
)

// String returns the string representation of the Style.
func (s Style) String() string {
	return string(s)
}

// NewStyleFromString converts a string to a Style, returning an error if the format is unsupported.
func NewStyleFromString(s string) (Style, error) {
	style := Style(s)
	switch style {
	case Stylish, Plain, JSON:
		return style, nil
	default:
		return "", fmt.Errorf("unable to parse style: unsupported format type %s", s)
	}
}

type diffFormatter interface {
	Render([]diff.Record) (string, error)
}

// New creates a formatter for the given format type.
func New(style Style) diffFormatter {
	switch style {
	case Stylish:
		return &stylishDiffFormatter{}
	case Plain:
		return &plainDiffFormatter{}
	case JSON:
		return &jsonDiffFormatter{}
	}

	panic(fmt.Sprintf("unsupported format type: %s", style))
}
