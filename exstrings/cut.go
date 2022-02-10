package exstrings

import "strings"

// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
//
// Port from Go 1.18
// https://cs.opensource.google/go/go/+/refs/tags/go1.18beta1:src/strings/strings.go;l=1177
func Cut(s, sep string) (before, after string, found bool) {
	if i := strings.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}
