package extime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMonotonic(t *testing.T) {
	assert.NotEqual(t, int64(0), Monotonic())
}

func TestMonoSub(t *testing.T) {
	mono1 := Monotonic()
	time.Sleep(1)
	mono2 := Monotonic()

	assert.Equal(t, time.Duration(mono1-mono2), MonoSub(mono1, mono2))
}
