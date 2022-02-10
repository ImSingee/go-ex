package set

type (
	Set struct {
		hash map[interface{}]nothing
	}

	nothing struct{}
)

// New Create a new set
func New(initial ...interface{}) *Set {
	s := &Set{make(map[interface{}]nothing)}

	for _, v := range initial {
		s.Insert(v)
	}

	return s
}

// Difference Find the difference between two sets
// 返回的是自己存在而传入的 set 不存在的元素集合
func (s *Set) Difference(set *Set) *Set {
	n := make(map[interface{}]nothing)

	for k := range s.hash {
		if _, exists := set.hash[k]; !exists {
			n[k] = nothing{}
		}
	}

	return &Set{n}
}

func (s *Set) All() []interface{} {
	if s == nil || len(s.hash) == 0 {
		return []interface{}{}
	}

	all := make([]interface{}, 0, len(s.hash))
	for i := range s.hash {
		all = append(all, i)
	}
	return all
}

// Do Call f for each item in the set
func (s *Set) Do(f func(interface{})) {
	for k := range s.hash {
		f(k)
	}
}

// DoE Call f for each item in the set
func (s *Set) DoE(f func(interface{}) error) error {
	for k := range s.hash {
		err := f(k)
		if err != nil {
			return err
		}
	}

	return nil
}

// Has Test to see whether or not the element is in the set
func (s *Set) Has(element interface{}) bool {
	_, exists := s.hash[element]
	return exists
}

// Insert Add element(s) to the set
func (s *Set) Insert(elements ...interface{}) {
	for _, e := range elements {
		s.hash[e] = nothing{}
	}
}

// Intersection Find the intersection of two sets
func (s *Set) Intersection(set *Set) *Set {
	n := make(map[interface{}]nothing)

	for k := range s.hash {
		if _, exists := set.hash[k]; exists {
			n[k] = nothing{}
		}
	}

	return &Set{n}
}

// Len Return the number of items in the set
func (s *Set) Len() int {
	return len(s.hash)
}

// ProperSubsetOf Test whether or not this set is a proper subset of "set"
func (s *Set) ProperSubsetOf(set *Set) bool {
	return s.SubsetOf(set) && s.Len() < set.Len()
}

// Remove an element from the set
func (s *Set) Remove(element interface{}) {
	delete(s.hash, element)
}

// SubsetOf Test whether or not this set is a subset of "set"
func (s *Set) SubsetOf(set *Set) bool {
	if s.Len() > set.Len() {
		return false
	}
	for k := range s.hash {
		if _, exists := set.hash[k]; !exists {
			return false
		}
	}
	return true
}

// Union Find the union of two sets
func (s *Set) Union(set *Set) *Set {
	n := make(map[interface{}]nothing)

	for k := range s.hash {
		n[k] = nothing{}
	}
	for k := range set.hash {
		n[k] = nothing{}
	}

	return &Set{n}
}

// Intersection 返回若干个 set 的交集
func Intersection(sets ...*Set) *Set {
	if len(sets) == 0 {
		return New()
	}

	b := sets[0]
	for _, s := range sets[1:] {
		b = b.Intersection(s)
	}

	return b
}
