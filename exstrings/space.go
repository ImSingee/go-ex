package exstrings

import (
	"regexp"
)

// FoldSpace replace any spaces (more than two) to one space
func FoldSpace(s string) string {
	return regexp.MustCompile("\\s{2,}").ReplaceAllString(s, " ")
}
