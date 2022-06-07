package simplelru

func (c *LRU[K, V]) zeroKey() K {
	return zero[K]()
}

func (c *LRU[K, V]) zeroValue() V {
	return zero[V]()
}

// helper for generic default value
func zero[T any]() T {
	var result T
	return result
}
