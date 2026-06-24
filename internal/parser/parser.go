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

	return normalize(parsed).(map[string]any), nil
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

func normalize(value any) any {
	switch typedValue := value.(type) {
	case map[string]any:
		result := make(map[string]any, len(typedValue))
		for key, nestedValue := range typedValue {
			result[key] = normalize(nestedValue)
		}

		return result

	case []any:
		result := make([]any, len(typedValue))
		for index, nestedValue := range typedValue {
			result[index] = normalize(nestedValue)
		}

		return result

	case int:
		return float64(typedValue)

	case int64:
		return float64(typedValue)

	case float32:
		return float64(typedValue)

	default:
		return value
	}
}
