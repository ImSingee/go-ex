package exlist

func In[T comparable](list []T, v T) bool {
	for _, vv := range list {
		if vv == v {
			return true
		}
	}
	return false
}
