package formatter

import (
	"code/internal/diff"
	"encoding/json"
	"fmt"
)

type jsonDiffFormatter struct{}

func (j *jsonDiffFormatter) Render(records []diff.Record) string {
	data := j.buildObject(records)

	result, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err) // TODO: should I handle properly? I think this is an unexpected error, so panic is ok.
	}

	return string(result)
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
			"type":  "added",
			"value": record.NewValue,
		}

	case diff.Removed:
		return map[string]any{
			"type":  "removed",
			"value": record.OldValue,
		}

	case diff.Changed:
		return map[string]any{
			"type": "changed",
			"old":  record.OldValue,
			"new":  record.NewValue,
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
