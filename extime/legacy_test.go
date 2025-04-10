//go:build !go1.23 || allow_linkname

package extime

import (
	"github.com/ImSingee/tt"
	"testing"
	"time"
)

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
