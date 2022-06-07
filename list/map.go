package list

func Map[K comparable, V any](l *List[K], f func(K) V) []V {
	if l == nil {
		return nil
	}

	result := make([]V, len(l.E))
	for i, s := range l.E {
		result[i] = f(s)
	}

	return result
}

func MapE[K comparable, V any](l *List[K], f func(K) (V, error)) ([]V, error) {
	if l == nil {
		return nil, nil
	}

	result := make([]V, len(l.E))
	var err error
	for i, s := range l.E {
		result[i], err = f(s)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
