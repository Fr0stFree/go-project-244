// Package diff builds an internal representation of differences between parsed data.
package diff

import (
	"reflect"

	"github.com/samber/lo"
)

type RecordState int

const (
	Added     RecordState = 0
	Removed   RecordState = 1
	Unchanged RecordState = 2
	Changed   RecordState = 3
)

type Record struct {
	Key      string
	State    RecordState
	OldValue any
	NewValue any
}

// Build compares two parsed configuration files and returns their difference.
func Build(left, right map[string]any) []Record {
	keys := lo.Union(lo.Keys(left), lo.Keys(right))

	nodes := make([]Record, 0, len(keys))
	for _, key := range keys {
		leftValue, isExistInLeft := left[key]
		rightValue, isExistInRight := right[key]
		node := createNode(key, leftValue, rightValue, isExistInLeft, isExistInRight)
		nodes = append(nodes, node)
	}

	return nodes
}

func createNode(key string, oldValue, newValue any, isExistBefore, isExistAfter bool) Record {
	if !isExistAfter && isExistBefore {
		return Record{key, Removed, oldValue, nil}
	}

	if isExistAfter && !isExistBefore {
		return Record{key, Added, nil, newValue}
	}

	if reflect.DeepEqual(oldValue, newValue) {
		return Record{key, Unchanged, oldValue, newValue}
	}

	return Record{key, Changed, oldValue, newValue}
}
