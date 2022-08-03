package set

import "sync"

// SyncSet provide nearly same methods as Set, but it is thread-safe
type SyncSet[T comparable] struct {
	mu sync.RWMutex
	in *Set[T]
}

// NewSync Create a new SyncSet
func NewSync[T comparable](initial ...T) *SyncSet[T] {
	s := New(initial...)

	return &SyncSet[T]{
		in: s,
	}
}

// NewSyncWithCapacity Create an empty SyncSet with specific capacity
func NewSyncWithCapacity[T comparable](capacity int) *SyncSet[T] {
	s := WithCapacity[T](capacity)

	return &SyncSet[T]{
		in: s,
	}
}

func (s *SyncSet[T]) rw(f func(*Set[T])) {
	s.mu.Lock()
	defer s.mu.Unlock()

	f(s.in)
}

func (s *SyncSet[T]) r(f func(*Set[T])) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	f(s.in)
}

// Has Test to see whether the element is in the set
func (s *SyncSet[T]) Has(element T) (exists bool) {
	s.r(func(in *Set[T]) {
		exists = in.Has(element)
	})
	return
}

func (s *SyncSet[T]) All() (all []T) {
	s.r(func(in *Set[T]) {
		all = in.All()
	})
	return
}

// Add element(s) to the set
func (s *SyncSet[T]) Add(elements ...T) {
	s.rw(func(in *Set[T]) {
		in.InsertAll(elements)
	})
}

// AddAll add elements to the set
func (s *SyncSet[T]) AddAll(elements []T) {
	s.rw(func(in *Set[T]) {
		in.InsertAll(elements)
	})
}

// Len Return the number of items in the set
func (s *SyncSet[T]) Len() (size int) {
	s.r(func(in *Set[T]) {
		size = in.Len()
	})
	return
}

// Size Return the number of items in the set
func (s *SyncSet[T]) Size() (size int) {
	s.r(func(in *Set[T]) {
		size = in.Len()
	})
	return
}

// Remove an element from the set
func (s *SyncSet[T]) Remove(element T) {
	s.rw(func(in *Set[T]) {
		in.Remove(element)
	})
}

// ContainsOrAdd add element if not exist
func (s *SyncSet[T]) ContainsOrAdd(element T) (added bool) {
	s.rw(func(in *Set[T]) {
		added = in.ContainsOrAdd(element)
	})
	return
}
