package ee

import (
	"fmt"
	"io"
)

func WithStackSkip(err error, skip int) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(StackTracer); ok {
		return err
	}

	return &withStack{
		err,
		callersSkip(skip),
	}
}

// WithStack annotates err with a stack trace at the point WithStack was called.
// If err is nil, WithStack returns nil.
// If err is already a StackTracer returns original error
func WithStack(err error) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(StackTracer); ok {
		return err
	}

	return &withStack{
		err,
		callers(0),
	}
}

type withStack struct {
	error
	*stack
}

func (w *withStack) Unwrap() error { return w.error }

func (w *withStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%+v", w.error)
			w.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, w.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", w.Error())
	}
}
