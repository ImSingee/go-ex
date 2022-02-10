package set

func NewFromString(list []string) *Set {
	s := &Set{make(map[interface{}]nothing)}

	for _, v := range list {
		s.Insert(v)
	}

	return s
}

func (s *Set) ToStringList() ([]string, bool) {
	result := make([]string, 0, s.Len())

	for k := range s.hash {
		ks, ok := k.(string)
		if !ok {
			return nil, false
		}

		result = append(result, ks)
	}

	return result, true
}
