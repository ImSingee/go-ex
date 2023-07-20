package mr

func Reduce[I, O any](arr []I, fn func(acc O, in I, index int) O, acc O) O {
	for i, v := range arr {
		acc = fn(acc, v, i)
	}
	return acc
}

func ReduceE[I, O any](arr []I, fn func(acc O, in I, index int) (O, error), acc O) (O, error) {
	for i, v := range arr {
		r, err := fn(acc, v, i)
		if err != nil {
			return acc, err
		}
		acc = r
	}
	return acc, nil
}
