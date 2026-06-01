// Package diff builds an internal representation of differences between parsed data.
package diff

import (
	"reflect"

	"github.com/samber/lo"
)

type NodeState int

const (
	Added     NodeState = 0
	Removed   NodeState = 1
	Unchanged NodeState = 2
	Changed   NodeState = 3
)

type Node struct {
	Key      string
	State    NodeState
	OldValue any
	NewValue any
}

// Build compares two parsed configuration files and returns their difference.
func Build(left, right map[string]any) []Node {
	keys := lo.Union(lo.Keys(left), lo.Keys(right))

	nodes := make([]Node, 0, len(keys))
	for _, key := range keys {
		leftValue, isExistInLeft := left[key]
		rightValue, isExistInRight := right[key]
		node := createNode(key, leftValue, rightValue, isExistInLeft, isExistInRight)
		nodes = append(nodes, node)
	}

	return nodes
}

func createNode(key string, oldValue, newValue any, isExistBefore, isExistAfter bool) Node {
	if !isExistAfter && isExistBefore {
		return Node{key, Removed, oldValue, nil}
	}

	if isExistAfter && !isExistBefore {
		return Node{key, Added, nil, newValue}
	}

	if reflect.DeepEqual(oldValue, newValue) {
		return Node{key, Unchanged, oldValue, newValue}
	}

	return Node{key, Changed, oldValue, newValue}
}
