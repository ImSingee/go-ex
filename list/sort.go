package list

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func (l *List[T]) Sort(less func(a, b T) bool) {
	SortFunc(l, less)
}

func (l *List[T]) SortStable(less func(a, b T) bool) {
	SortStableFunc(l, less)
}

func Sort[T constraints.Ordered](l *List[T]) {
	slices.Sort(l.E)
}

func SortFunc[T comparable](l *List[T], less func(a, b T) bool) {
	slices.SortFunc(l.E, less)
}

func SortStableFunc[T comparable](l *List[T], less func(a, b T) bool) {
	slices.SortStableFunc(l.E, less)
}
