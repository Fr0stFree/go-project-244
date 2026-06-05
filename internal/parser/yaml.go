package parser

import "gopkg.in/yaml.v3"

type yamlParser struct{}

func (y yamlParser) run(payload []byte) (map[string]any, error) {
	var result map[string]any

	err := yaml.Unmarshal(payload, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}
