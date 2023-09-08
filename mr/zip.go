package mr

// Zip joins multiple slices
//
// For example, Zip([a, b, c], [d, e, f]) returns [[a, d], [b, e], [c, f]]
// Each slice must have the same length (or it's an UB or even panic)
func Zip[T any](args ...[]T) [][]T {
	if len(args) == 0 {
		return nil
	}

	length := len(args[0])
	result := make([][]T, length)

	for i := 0; i < length; i++ {
		result[i] = make([]T, len(args))
		for j, arg := range args {
			result[i][j] = arg[i]
		}
	}

	return result
}

// ZipFlat likes Zip, but returns a flat slice
func ZipFlat[T any](args ...[]T) []T {
	return Flat(Zip(args...))
}
