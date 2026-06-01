// Package diff builds an internal representation of differences between parsed data.
package diff

// Build compares two parsed configuration files and returns their difference.
func Build(left, right map[string]any) map[string]any {
	_ = left
	_ = right
	result := make(map[string]any)

	return result
}
