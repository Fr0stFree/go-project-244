package formatter

import (
	"code/internal/diff"
	"encoding/json"
	"fmt"
)

type jsonDiffFormatter struct{}

func (j *jsonDiffFormatter) Render(records []diff.Record) (string, error) {
	result := make(map[string]diff.Record, len(records))

	for _, record := range records {
		result[record.Key] = record
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON diff: %w", err)
	}

	return string(data), nil
}
