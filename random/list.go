package random

import "math/rand"

func defaultValue[T any]() (x T) {
	return
}

func Choice[T any](l []T) T {
	n := len(l)
	if n == 0 {
		return defaultValue[T]()
	}

	return l[rand.Intn(n)]
}
