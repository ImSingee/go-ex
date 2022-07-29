package ee

import "fmt"

func Panic(message string) {
	panic(&fundamental{
		msg:   message,
		stack: callers(0),
	})
}

func Panicf(format string, args ...interface{}) {
	panic(&fundamental{
		msg:   fmt.Sprintf(format, args...),
		stack: callers(0),
	})
}

func PanicE(err error) {
	if _, ok := err.(StackTracer); ok {
		panic(err)
	}

	panic(&withStack{
		err,
		callers(0),
	})
}

func PanicWrap(err error, message string) {
	panic(wrap(err, message))
}

func PanicWrapf(err error, format string, args ...interface{}) {
	panic(wrap(err, fmt.Sprintf(format, args...)))
}

func ErrFromPanic(v interface{}) error {
	err, _ := v.(stacker)
	return err
}
