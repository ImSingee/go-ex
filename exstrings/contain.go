package exstrings

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
