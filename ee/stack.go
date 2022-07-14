package ee

import (
	"fmt"
	"runtime"
	"strings"
)

// stack represents a stack of program counters.
type stack []uintptr

func (s *stack) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case st.Flag('+'):
			for _, pc := range *s {
				f := Frame(pc)
				_, _ = fmt.Fprintf(st, "\n%+v", f)
			}
		}
	}
}

func (s *stack) StackTrace() StackTrace {
	f := make([]Frame, len(*s))
	for i := 0; i < len(f); i++ {
		f[i] = Frame((*s)[i])
	}
	return f
}

func (s *stack) singeeErrStack() *stack {
	return s
}

func callersSkip(skip int) *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3+skip, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

func callers(skip int) *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(skip+3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

// funcname removes the path prefix component of a function's name reported by func.Name().
func funcname(name string) string {
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}

type StackTracer interface {
	StackTrace() StackTrace
}

type stacker interface {
	error
	singeeErrStack() *stack
}
