package formatter

import (
	"code/internal/diff"
	"encoding/json"
	"fmt"
)

type jsonDiffFormatter struct{}

func (j *jsonDiffFormatter) Render(records []diff.Record) (string, error) {
	result, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}

	return string(result), nil
}
