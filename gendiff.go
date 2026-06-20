// Package code provides functionality for comparing configuration files.
package code

import (
	"code/internal/diff"
	"code/internal/formatter"
	"code/internal/parser"
)

// GenDiff compares two configuration files and returns their formatted difference.
func GenDiff(firstPath, secondPath, outputFormat string) (string, error) {
	style, err := formatter.NewStyleFromString(outputFormat)
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

	formatter := formatter.New(style)

	output, err := formatter.Render(difference)
	if err != nil {
		return "", err
	}

	return output, nil
}
