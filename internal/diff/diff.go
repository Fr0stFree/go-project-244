// Package diff builds an internal representation of differences between parsed data.
package diff

import (
	"cmp"
	"reflect"
	"slices"

	"github.com/samber/lo"
)

type recordState string

// RecordState represents the type of difference between two values.
const (
	Added     recordState = "added"
	Removed   recordState = "removed"
	Unchanged recordState = "unchanged"
	Changed   recordState = "changed"
	Nested    recordState = "nested"
)

// Record represents a difference between two values in a map.
type Record struct {
	Key      string      `json:"key"`
	State    recordState `json:"type"`
	OldValue any         `json:"old,omitempty"`
	NewValue any         `json:"new,omitempty"`
	Children []Record    `json:"children,omitempty"`
}

// Build compares two parsed configuration files and returns their difference.
func Build(left, right map[string]any) []Record {
	keys := lo.Union(lo.Keys(left), lo.Keys(right))

	records := make([]Record, 0, len(keys))
	for _, key := range keys {
		leftValue, existsInLeft := left[key]
		rightValue, existsInRight := right[key]
		record := newRecord(key, leftValue, rightValue, existsInLeft, existsInRight)
		records = append(records, record)
	}

	slices.SortFunc(records, func(a, b Record) int {
		return cmp.Compare(a.Key, b.Key)
	})

	return records
}

func newRecord(key string, oldValue, newValue any, existsBefore, existsAfter bool) Record {
	oldMap, wasMap := oldValue.(map[string]any)
	newMap, isMap := newValue.(map[string]any)

	switch {
	case !existsAfter && existsBefore:
		return Record{Key: key, State: Removed, OldValue: oldValue, NewValue: nil}

	case existsAfter && !existsBefore:
		return Record{Key: key, State: Added, OldValue: nil, NewValue: newValue}

	case wasMap && isMap:
		return Record{Key: key, State: Nested, Children: Build(oldMap, newMap)}

	case reflect.DeepEqual(oldValue, newValue):
		return Record{Key: key, State: Unchanged, OldValue: oldValue, NewValue: newValue}

	default:
		return Record{Key: key, State: Changed, OldValue: oldValue, NewValue: newValue}
	}
}
