package excache

import (
	"fmt"
	"github.com/ImSingee/go-ex/ee"
	"sync"
	"time"
)

type CachedValue[T any] struct {
	l sync.Mutex
	v T

	nextFetch     time.Time
	fetchInterval time.Duration
	fetchFn       func() (T, error)
}

var ErrFetchFnNotSet = fmt.Errorf("fetch function not set")

func NewCachedValue[T any](fetchInterval time.Duration, fetchFn ...func() (T, error)) *CachedValue[T] {
	cv := &CachedValue[T]{
		fetchInterval: fetchInterval,
	}

	if len(fetchFn) > 0 {
		cv.fetchFn = fetchFn[0]
	}

	return cv
}

func (cv *CachedValue[T]) Get(fetchFn ...func() (T, error)) (T, error) {
	var fn func() (T, error)
	if len(fetchFn) > 0 {
		fn = fetchFn[0]
	} else {
		fn = cv.fetchFn
	}

	if fn == nil {
		return nil, ee.WithStack(ErrFetchFnNotSet)
	}

	cv.l.Lock()
	defer cv.l.Unlock()

	if time.Now().After(cv.nextFetch) {
		v, err := fn()
		if err != nil {
			return cv.v, ee.WithStack(err)
		}

		cv.v = v
		cv.nextFetch = time.Now().Add(cv.fetchInterval)
	}

	return cv.v, nil
}

func (cv *CachedValue[T]) Refresh(fetchFn ...func() (T, error)) (T, error) {
	var fn func() (T, error)
	if len(fetchFn) > 0 {
		fn = fetchFn[0]
	} else {
		fn = cv.fetchFn
	}

	if fn == nil {
		return nil, ee.WithStack(ErrFetchFnNotSet)
	}

	cv.l.Lock()
	defer cv.l.Unlock()

	v, err := fn()
	if err != nil {
		return cv.v, ee.WithStack(err)
	}

	cv.v = v
	cv.nextFetch = time.Now().Add(cv.fetchInterval)

	return cv.v, nil
}

// refresh immediately
// must acquire lock before call (and release it after)
func (cv *CachedValue[T]) refresh(f func() (T, error)) (T, error) {
	v, err := f()
	if err != nil {
		return v, err
	}

	cv.v = v
	cv.nextFetch = time.Now().Add(cv.fetchInterval)

	return v, nil
}

// SetFetchInterval change fetch interval (will not affect nextFetch time)
func (cv *CachedValue[T]) SetFetchInterval(fetchInterval time.Duration) {
	cv.fetchInterval = fetchInterval
}

// SetNextFetchTime change next fetch time
func (cv *CachedValue[T]) SetNextFetchTime(t time.Time) {
	cv.nextFetch = t
}
