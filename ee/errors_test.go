package ee

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func f1() error {
	return New("from f1")
}

func f2() error {
	return Wrap(f1(), "from f2")
}

func TestErrorStack(t *testing.T) {
	err := f2().(stacker)

	stack := err.singeeErrStack()
	s := fmt.Sprintf("%+v\n", stack)
	assert.True(t, strings.Contains(s, "ee.f1"))
}
