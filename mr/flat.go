package mr

func Flat[T any](arr [][]T) []T {
	return Reduce(arr, func(acc []T, in []T, index int) []T {
		return append(acc, in...)
	}, nil)
}
