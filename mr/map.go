package mr

func Map[I, O any](arr []I, fn func(in I, index int) O) []O {
	result := make([]O, len(arr))
	for i, v := range arr {
		result[i] = fn(v, i)
	}
	return result
}

func MapE[I, O any](arr []I, fn func(in I, index int) (O, error)) ([]O, error) {
	result := make([]O, len(arr))
	for i, v := range arr {
		r, err := fn(v, i)
		if err != nil {
			return nil, err
		}
		result[i] = r
	}
	return result, nil
}
