package utils

import (
	"strings"
)

// CommentLines prefixes each line in lines with "// ".
func CommentLines(lines string) string {
	lines = "// " + strings.Replace(lines, "\n", "\n// ", -1)
	return lines
}
