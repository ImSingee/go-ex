package mr

func zero[T any]() (v T) {
	return
}

// Find returns the first element in the slice that satisfies the provided testing function. Otherwise zero value is returned.
func Find[T any](arr []T, predict func(T) bool) T {
	for _, v := range arr {
		if predict(v) {
			return v
		}
	}
	return zero[T]()
}

// FindP returns the first element in the slice that satisfies the provided testing function. Otherwise nil is returned.
func FindP[T any](arr []T, predict func(*T) bool) *T {
	for _, v := range arr {
		if predict(&v) {
			return &v
		}
	}
	return nil
}

// FindIndex returns the index of the first element in the slice that satisfies the provided testing function. Otherwise -1 is returned.
func FindIndex[T any](arr []T, predict func(T) bool) int {
	for i, v := range arr {
		if predict(v) {
			return i
		}
	}
	return -1
}
