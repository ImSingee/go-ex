package exstrings

import "strings"

// Capitalize convert first letter to upper and others remain
func Capitalize(s string) string {
	if s == "" {
		return s
	}

	return strings.ToUpper(s[:1]) + s[1:]
}
