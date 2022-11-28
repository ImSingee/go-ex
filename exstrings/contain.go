package exstrings

import "strings"

func ContainsOnly(s, chars string) bool {
	if len(s) == 0 || len(chars) == 0 {
		return false
	}
outer:
	for _, r := range s {
		for _, c := range chars {
			if r == c {
				continue outer
			}
		}
		return false
	}
	return true
}

// ContainsAny returns true if s contains any specified substring
func ContainsAny(s string, ss ...string) bool {
	for _, sss := range ss {
		if strings.Contains(s, sss) {
			return true
		}
	}
	return false
}
