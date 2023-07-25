package mr

func zero[T any]() (v T) {
	return
}

func Find[T any](arr []T, predict func(T) bool) T {
	for _, v := range arr {
		if predict(v) {
			return v
		}
	}
	return zero[T]()
}

func FindP[T any](arr []T, predict func(*T) bool) *T {
	for _, v := range arr {
		if predict(&v) {
			return &v
		}
	}
	return nil
}
