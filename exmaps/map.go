package exmaps

func MapFromKeys(keys []string) map[string]struct{} {
	m := make(map[string]struct{}, len(keys))

	for _, key := range keys {
		m[key] = struct{}{}
	}

	return m
}

func KeyMapContains(keysMap map[string]struct{}, key string) bool {
	_, ok := keysMap[key]
	return ok
}
