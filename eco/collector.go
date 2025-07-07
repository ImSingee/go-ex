//go:build go1.20

// Package eco is a package that provides a collector for errors
//
// This package uses the errors.Join function under the hood so it's compatible with Go 1.20 and above.
// If you are using Go 1.19 or below, you can see https://github.com/ImSingee/eco.
package eco

import (
	"errors"
	"fmt"
)

type Collector struct {
	Errors []error
}

func New() Collector {
	return Collector{}
}

func (e Collector) Err() error {
	return errors.Join(e.Errors...)
}

func (e Collector) Error() string {
	return e.Err().Error()
}

func (e Collector) Unwrap() []error {
	return e.Errors
}

func (e Collector) IsError() bool {
	return len(e.Errors) != 0
}

// Collect collect error if not nil
func (e *Collector) Collect(err error) {
	if err != nil {
		e.Errors = append(e.Errors, err)
	}
}

// C collect error if not nil and return received error
//
// C is same as Collect, but also return received error
func (e *Collector) C(err error) error {
	if err != nil {
		e.Errors = append(e.Errors, err)
	}

	return err
}

// Do execute the function and collect error if not nil
func (e *Collector) Do(f func() error) {
	e.Collect(e.do(f))
}

// D collect error if not nil and return received error
//
// D is same as Do, but also return received error
func (e *Collector) D(f func() error) (err error) {
	return e.C(e.do(f))
}

func (e *Collector) do(f func() error) (err error) {
	defer func() {
		v := recover()
		if v != nil {
			err = fmt.Errorf("panic: %v", v)
		}
	}()

	err = f()
	return
}

// Process execute the functions one by one and collect error if not nil
// When any error is collected (the process return error), the remaining processes are not executed
func (e *Collector) Process(processes ...func() error) error {
	for _, p := range processes {
		if e.IsError() {
			return e.Err()
		}

		e.Do(p)
	}

	return e.Err()
}

// Do execute the functions one by one and collect error if not nil
// When any error is collected (the process return error), the remaining processes are not executed
func Do(processes ...func() error) error {
	e := New()
	return e.Process(processes...)
}
