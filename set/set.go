package set

type (
	Set[T comparable] struct {
		hash map[T]nothing
	}

	nothing struct{}
)

// New Create a new set
// The capacity of created set is length of initial,
// if you want to customize it, please use WithCapacity and .Insert / .InsertAll
func New[T comparable](initial ...T) *Set[T] {
	s := &Set[T]{make(map[T]nothing, len(initial))}

	for _, v := range initial {
		s.Insert(v)
	}

	return s
}

// WithCapacity Create an empty set with specific capacity
func WithCapacity[T comparable](capacity int) *Set[T] {
	s := &Set[T]{make(map[T]nothing, capacity)}

	return s
}

// Difference Find the difference between two sets
// 返回的是自己存在而传入的 set 不存在的元素集合
func (s *Set[T]) Difference(another *Set[T]) *Set[T] {
	n := make(map[T]nothing)

	for k := range s.hash {
		if _, exists := another.hash[k]; !exists {
			n[k] = nothing{}
		}
	}

	return &Set[T]{n}
}

func (s *Set[T]) All() []T {
	if s == nil || len(s.hash) == 0 {
		return []T{}
	}

	all := make([]T, 0, len(s.hash))
	for i := range s.hash {
		all = append(all, i)
	}
	return all
}

// Do Call f for each item in the set
func (s *Set[T]) Do(f func(T)) {
	for k := range s.hash {
		f(k)
	}
}

// DoE Call f for each item in the set
func (s *Set[T]) DoE(f func(T) error) error {
	for k := range s.hash {
		err := f(k)
		if err != nil {
			return err
		}
	}

	return nil
}

// Contains Test to see whether the element is in the set
func (s *Set[T]) Contains(element T) bool {
	return s.Has(element)
}

// Has Test to see whether the element is in the set
func (s *Set[T]) Has(element T) bool {
	_, exists := s.hash[element]
	return exists
}

// Add element(s) to the set
func (s *Set[T]) Add(elements ...T) {
	s.InsertAll(elements)
}

// AddAll add elements to the set
func (s *Set[T]) AddAll(elements []T) {
	s.InsertAll(elements)
}

// Insert add element(s) to the set
func (s *Set[T]) Insert(elements ...T) {
	s.InsertAll(elements)
}

// InsertAll add elements to the set
func (s *Set[T]) InsertAll(elements []T) {
	for _, e := range elements {
		s.hash[e] = nothing{}
	}
}

// Intersection Find the intersection of two sets
func (s *Set[T]) Intersection(another *Set[T]) *Set[T] {
	n := make(map[T]nothing)

	for k := range s.hash {
		if _, exists := another.hash[k]; exists {
			n[k] = nothing{}
		}
	}

	return &Set[T]{n}
}

// Len Return the number of items in the set
func (s *Set[T]) Len() int {
	return len(s.hash)
}

// Size Return the number of items in the set
func (s *Set[T]) Size() int {
	return s.Len()
}

// ProperSubsetOf Test whether this set is a proper subset of "set"
func (s *Set[T]) ProperSubsetOf(another *Set[T]) bool {
	return s.SubsetOf(another) && s.Len() < another.Len()
}

// Remove an element from the set
func (s *Set[T]) Remove(element T) {
	delete(s.hash, element)
}

// SubsetOf Test whether this set is a subset of "set"
func (s *Set[T]) SubsetOf(set *Set[T]) bool {
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
func (s *Set[T]) Union(set *Set[T]) *Set[T] {
	n := make(map[T]nothing)

	for k := range s.hash {
		n[k] = nothing{}
	}
	for k := range set.hash {
		n[k] = nothing{}
	}

	return &Set[T]{n}
}

// ContainsOrAdd add element if not exist
func (s *Set[T]) ContainsOrAdd(element T) (added bool) {
	if s.Has(element) {
		return false
	}

	s.Add(element)
	return true
}

// Intersection 返回若干个 set 的交集
func Intersection[T comparable](sets ...*Set[T]) *Set[T] {
	if len(sets) == 0 {
		return New[T]()
	}

	b := sets[0]
	for _, s := range sets[1:] {
		b = b.Intersection(s)
	}

	return b
}
