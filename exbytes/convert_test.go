package exbytes

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	var s string
	{
		b := []byte("Hello world")
		s = ToString(b)
	}

	runtime.GC()

	assert.Equal(t, "Hello world", s)
}
