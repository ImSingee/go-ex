package exbytes

import (
	"github.com/ImSingee/tt"
	"runtime"
	"testing"
)

func TestToString(t *testing.T) {
	var s string
	{
		b := []byte("Hello world")
		s = ToString(b)
	}

	runtime.GC()

	tt.AssertEqual(t, "Hello world", s)
}
