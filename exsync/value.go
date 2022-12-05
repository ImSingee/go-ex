package exsync

import "sync"

type Value[T any] struct {
	sync.Mutex
	v T
}

// Set value (will acquire lock before set and unlock after set)
func (v *Value[T]) Set(newValue T) {
	v.Lock()
	defer v.Unlock()

	v.v = newValue
}

// Get value and acquire lock
// You must manually call Unlock() after using the value
func (v *Value[T]) Get() T {
	v.Lock()

	return v.v
}

// With will call f with the value and unlock it after f returns
func (v *Value[T]) With(f func(v T)) {
	v.Lock()
	defer v.Unlock()

	f(v.v)
}

// WithE equals With, but f can return an error
func (v *Value[T]) WithE(f func(v T) error) error {
	v.Lock()
	defer v.Unlock()

	return f(v.v)
}
