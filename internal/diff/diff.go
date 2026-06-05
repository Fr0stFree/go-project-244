// Package diff builds an internal representation of differences between parsed data.
package diff

import (
	"cmp"
	"reflect"
	"slices"

	"github.com/samber/lo"
)

type recordState int

// RecordState represents the type of difference between two values.
const (
	Added recordState = iota
	Removed
	Unchanged
	Changed
	Nested
)

// Record represents a difference between two values in a map.
type Record struct {
	Key      string
	State    recordState
	OldValue any
	NewValue any
	Children []Record
}

// Build compares two parsed configuration files and returns their difference.
func Build(left, right map[string]any) []Record {
	keys := lo.Union(lo.Keys(left), lo.Keys(right))

	records := make([]Record, 0, len(keys))
	for _, key := range keys {
		leftValue, isExistInLeft := left[key]
		rightValue, isExistInRight := right[key]
		record := newRecord(key, leftValue, rightValue, isExistInLeft, isExistInRight)
		records = append(records, record)
	}

	slices.SortFunc(records, func(a, b Record) int {
		return cmp.Compare(a.Key, b.Key)
	})

	return records
}

func newRecord(key string, oldValue, newValue any, isExistBefore, isExistAfter bool) Record {
	if !isExistAfter && isExistBefore {
		return Record{Key: key, State: Removed, OldValue: oldValue, NewValue: nil}
	}

	oldMap, isOldMap := oldValue.(map[string]any)
	newMap, isNewMap := newValue.(map[string]any)

	if isOldMap && isNewMap {
		return Record{Key: key, State: Nested, OldValue: nil, NewValue: nil, Children: Build(oldMap, newMap)}
	}

	if isExistAfter && !isExistBefore {
		return Record{Key: key, State: Added, OldValue: nil, NewValue: newValue}
	}

	if reflect.DeepEqual(oldValue, newValue) {
		return Record{Key: key, State: Unchanged, OldValue: oldValue, NewValue: newValue}
	}

	return Record{Key: key, State: Changed, OldValue: oldValue, NewValue: newValue}
}
