package formatter

type stylishDiffFormatter struct{}

// Render converts a diff into the stylish string representation.
func (s *stylishDiffFormatter) Render(map[string]any) string {
	return ""
}
