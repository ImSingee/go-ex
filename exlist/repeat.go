package exlist

// Repeat returns a slice containing n copies of item.
func Repeat[T any](n int, item T) []T {
	l := make([]T, n)

	for i := 0; i < n; i++ {
		l[i] = item
	}

	return l
}
