package queue

import (
	"bytes"
	"fmt"
)

// Queue represents a double-ended queue.
// The zero value is an empty queue ready to use.
type Queue[T any] struct {
	// PushBack writes to rep[back] then increments back; PushFront
	// decrements front then writes to rep[front]; len(rep) is a power
	// of two; unused slots are nil and not garbage.
	rep    []T
	front  int
	back   int
	length int
}

// New returns an initialized empty queue.
func New[T any](initialCapacity int) *Queue[T] {
	return &Queue[T]{
		rep: make([]T, 1, initialCapacity+1),
	}
}

// Len returns the number of elements of queue q.
func (q *Queue[T]) Len() int {
	return q.length
}

// empty returns true if the queue q has no elements.
func (q *Queue[T]) empty() bool {
	return q.length == 0
}

// full returns true if the queue q is at capacity.
func (q *Queue[T]) full() bool {
	return q.length == len(q.rep)
}

// sparse returns true if the queue q has excess capacity.
func (q *Queue[T]) sparse() bool {
	return 1 < q.length && q.length < len(q.rep)/4
}

// resize adjusts the size of queue q's underlying slice.
func (q *Queue[T]) resize(size int) {
	adjusted := make([]T, size)
	if q.front < q.back {
		// rep not "wrapped" around, one copy suffices
		copy(adjusted, q.rep[q.front:q.back])
	} else {
		// rep is "wrapped" around, need two copies
		n := copy(adjusted, q.rep[q.front:])
		copy(adjusted[n:], q.rep[:q.back])
	}
	q.rep = adjusted
	q.front = 0
	q.back = q.length
}

// lazyGrow grows the underlying slice if necessary.
func (q *Queue[T]) lazyGrow() {
	if q.full() {
		q.resize(len(q.rep) * 2)
	}
}

// lazyShrink shrinks the underlying slice if advisable.
func (q *Queue[T]) lazyShrink() {
	if q.sparse() {
		q.resize(len(q.rep) / 2)
	}
}

// String returns a string representation of queue q formatted
// from front to back.
func (q *Queue[T]) String() string {
	var result bytes.Buffer
	result.WriteByte('[')
	j := q.front
	for i := 0; i < q.length; i++ {
		result.WriteString(fmt.Sprintf("%v", q.rep[j]))
		if i < q.length-1 {
			result.WriteByte(' ')
		}
		j = q.inc(j)
	}
	result.WriteByte(']')
	return result.String()
}

// inc returns the next integer position wrapping around queue q.
func (q *Queue[T]) inc(i int) int {
	return (i + 1) & (len(q.rep) - 1) // requires l = 2^n
}

// dec returns the previous integer position wrapping around queue q.
func (q *Queue[T]) dec(i int) int {
	return (i - 1) & (len(q.rep) - 1) // requires l = 2^n
}

// Front returns the first element of queue q or nil.
func (q *Queue[T]) Front() T {
	// no need to check q.empty(), unused slots are nil
	return q.rep[q.front]
}

// Back returns the last element of queue q or nil.
func (q *Queue[T]) Back() T {
	// no need to check q.empty(), unused slots are nil
	return q.rep[q.dec(q.back)]
}

// PushFront inserts a new value v at the front of queue q.
func (q *Queue[T]) PushFront(v T) {
	q.lazyGrow()
	q.front = q.dec(q.front)
	q.rep[q.front] = v
	q.length++
}

// PushBack inserts a new value v at the back of queue q.
func (q *Queue[T]) PushBack(v T) {
	q.lazyGrow()
	q.rep[q.back] = v
	q.back = q.inc(q.back)
	q.length++
}

// PopFront removes and returns the first element of queue q or nil.
func (q *Queue[T]) PopFront() T {
	if q.empty() {
		return nilValue[T]()
	}
	v := q.rep[q.front]
	q.rep[q.front] = nilValue[T]() // unused slots must be nil
	q.front = q.inc(q.front)
	q.length--
	q.lazyShrink()
	return v
}

// PopBack removes and returns the last element of queue q or nil.
func (q *Queue[T]) PopBack() T {
	if q.empty() {
		return nilValue[T]()
	}
	q.back = q.dec(q.back)
	v := q.rep[q.back]
	q.rep[q.back] = nilValue[T]() // unused slots must be nil
	q.length--
	q.lazyShrink()
	return v
}

func nilValue[T any]() (v T) {
	return
}
