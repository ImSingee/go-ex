package exstrings

import "strings"

func FirstLine(s string) string {
	s = strings.TrimSpace(s)

	end := strings.Index(s, "\n")
	if end == -1 { // no new line, return entire string
		return s
	} else {
		return strings.TrimSpace(s[:end])
	}
}
