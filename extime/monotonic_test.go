package extime

import (
	"github.com/ImSingee/tt"
	"testing"
	"time"
)

func TestMonotonic(t *testing.T) {
	tt.AssertNotEqual(t, int64(0), Monotonic())
}

func TestStartNano(t *testing.T) {
	tt.AssertNotEqual(t, int64(0), startNano)
}

func TestGetTimeMono(t *testing.T) {
	now := time.Now()

	tt.AssertNotEqual(t, int64(0), GetTimeMono(now))

	time.Sleep(1)
	now2 := time.Now()
	tt.AssertTrue(t, GetTimeMono(now2) > GetTimeMono(now))

	time.Sleep(1)
	nano3 := Monotonic()
	tt.AssertTrue(t, nano3 > GetTimeMono(now2))
}

func TestMonoSub(t *testing.T) {
	mono1 := Monotonic()
	time.Sleep(1)
	mono2 := Monotonic()

	tt.AssertEqual(t, time.Duration(mono1-mono2), MonoSub(mono1, mono2))
}
