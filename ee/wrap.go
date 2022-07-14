package ee

import (
	"fmt"
)

func wrap(err error, message string) error {
	var cs *stack
	if e, ok := err.(stacker); ok {
		cs = e.singeeErrStack()
	} else {
		cs = callers(1)
	}

	err = &withMessage{
		cause: err,
		msg:   message,
	}
	return &withStack{
		err,
		cs,
	}
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message.
// If err is nil, Wrap returns nil.
// If err contains a stack trace already, will use original stack trace
func Wrap(err error, message string) error {
	return wrap(err, message)
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf is called, and the format specifier.
// If err is nil, Wrapf returns nil.
// If err contains a stack trace already, will use original stack trace
func Wrapf(err error, format string, args ...interface{}) error {
	return wrap(err, fmt.Sprintf(format, args...))
}

func WrapIfError(err error, message string) error {
	if err == nil {
		return nil
	}

	return wrap(err, message)
}

func WrapfIfError(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return wrap(err, fmt.Sprintf(format, args...))
}
