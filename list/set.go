package list

import "github.com/ImSingee/go-ex/set"

func FromSet[T comparable](s *set.Set[T]) *List[T] {
	return &List[T]{s.All()}
}

func (l *List[T]) ToSet() *set.Set[T] {
	return set.New(l.E...)
}
