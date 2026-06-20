package formatter

import (
	"code/internal/diff"
	"encoding/json"
	"fmt"
)

type jsonDiffFormatter struct{}

const (
	jsonNodeType        = "type"
	jsonNodeValue       = "value"
	jsonNodeTypeAdded   = "added"
	jsonNodeTypeRemoved = "removed"
	jsonNodeTypeChanged = "changed"
	jsonNodeNewValue    = "new"
	jsonNodeOldValue    = "old"
)

func (j *jsonDiffFormatter) Render(records []diff.Record) (string, error) {
	data := j.buildObject(records)

	result, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}

	return string(result), nil
}

func (j *jsonDiffFormatter) buildObject(records []diff.Record) map[string]any {
	result := make(map[string]any)

	for _, record := range records {
		node := j.buildNode(record)
		if node == nil {
			continue
		}

		result[record.Key] = node
	}

	return result
}

func (j *jsonDiffFormatter) buildNode(record diff.Record) any {
	switch record.State {
	case diff.Unchanged:
		return nil

	case diff.Added:
		return map[string]any{
			jsonNodeType:  jsonNodeTypeAdded,
			jsonNodeValue: record.NewValue,
		}

	case diff.Removed:
		return map[string]any{
			jsonNodeType:  jsonNodeTypeRemoved,
			jsonNodeValue: record.OldValue,
		}

	case diff.Changed:
		return map[string]any{
			jsonNodeType:     jsonNodeTypeChanged,
			jsonNodeOldValue: record.OldValue,
			jsonNodeNewValue: record.NewValue,
		}

	case diff.Nested:
		children := j.buildObject(record.Children)
		if len(children) == 0 {
			return nil
		}

		return children

	default:
		panic(fmt.Sprintf("unknown diff record state: %v", record.State))
	}
}
