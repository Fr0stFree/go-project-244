// Package code provides functionality for comparing configuration files.
package code

import (
	"code/internal/diff"
	"code/internal/formatter"
	"code/internal/parser"
)

// GenDiff compares two configuration files and returns their formatted difference.
func GenDiff(firstPath, secondPath, format string) (string, error) {
	formatter, err := formatter.New(format)
	if err != nil {
		return "", err
	}

	left, err := parser.ParseFile(firstPath)
	if err != nil {
		return "", err
	}

	right, err := parser.ParseFile(secondPath)
	if err != nil {
		return "", err
	}

	difference := diff.Build(left, right)

	return formatter.Render(difference), nil
}
