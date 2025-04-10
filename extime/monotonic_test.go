package extime

import (
	"github.com/ImSingee/tt"
	"testing"
	"time"
)

func TestMonotonic(t *testing.T) {
	tt.AssertNotEqual(t, int64(0), Monotonic())
}

func TestMonoSub(t *testing.T) {
	mono1 := Monotonic()
	time.Sleep(1)
	mono2 := Monotonic()

	tt.AssertEqual(t, time.Duration(mono1-mono2), MonoSub(mono1, mono2))
}
