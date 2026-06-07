// Package parser reads configuration files and parses them into maps.
package parser

import (
	"os"
	"path/filepath"
)

type parser interface {
	run([]byte) (map[string]any, error)
}

// ParseFile reads a file and parses its content based on the file extension.
func ParseFile(path string) (map[string]any, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, newParseError(err, path)
	}

	parser, err := selectParser(filepath.Ext(path))
	if err != nil {
		return nil, newParseError(err, path)
	}

	parsed, err := parser.run(raw)
	if err != nil {
		return nil, newParseError(err, path)
	}

	return parsed, nil
}

func selectParser(fileExt string) (parser, error) {
	switch fileExt {
	case ".json":
		return jsonParser{}, nil
	case ".yaml", ".yml":
		return yamlParser{}, nil
	case "":
		return nil, ErrNoFileExtension
	default:
		return nil, ErrUnsupportedFileType
	}
}
