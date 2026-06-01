package parser

import (
	"encoding/json"
)

type jsonParser struct{}

func (j jsonParser) run(payload []byte) (map[string]any, error) {
	var result map[string]any

	err := json.Unmarshal(payload, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
