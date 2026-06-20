// Package parser reads configuration files and parses them into maps.
package parser

import (
	"fmt"
	"os"
	"path/filepath"
)

type parseFunc func([]byte) (map[string]any, error)

// ParseFile reads a file and parses its content based on the file extension.
func ParseFile(path string) (map[string]any, error) {
	parseFunc, err := selectParseFunc(filepath.Ext(path))
	if err != nil {
		return nil, newParseError(err, path)
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	parsed, err := parseFunc(raw)
	if err != nil {
		return nil, newParseError(err, path)
	}

	return parsed, nil
}

func selectParseFunc(fileExt string) (parseFunc, error) {
	switch fileExt {
	case ".json":
		return parseJSON, nil
	case ".yaml", ".yml":
		return parseYAML, nil
	case "":
		return nil, ErrNoFileExtension
	default:
		return nil, ErrUnsupportedFileType
	}
}
