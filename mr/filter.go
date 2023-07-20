package mr

func Filter[T any](arr []T, fn func(in T, index int) bool) []T {
	result := make([]T, 0, len(arr))
	for i, v := range arr {
		if fn(v, i) {
			result = append(result, v)
		}
	}
	return result
}
