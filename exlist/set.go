package exlist

// Deprecated: Use set package instead
func ToSet(l []string) map[string]struct{} {
	m := make(map[string]struct{}, len(l))

	for _, e := range l {
		m[e] = struct{}{}
	}

	return m
}
