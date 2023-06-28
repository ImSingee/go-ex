package extime

import (
	"time"

	_ "unsafe"
)

//go:noescape
//go:linkname nanotime runtime.nanotime
func nanotime() int64

// Monotonic returns the system monotonic clock, reported in nanoseconds.
//
// Warning: though sometimes the return value looks like a Unix
// nanosecond timestamp, it is not. Only use it to compute time intervals.
func Monotonic() int64 {
	return nanotime()
}

//go:linkname startNano time.startNano
var startNano int64

//go:noescape
//go:linkname timeMono time.(*Time).mono
func timeMono(t *time.Time) int64

// GetTimeMono returns the mono part of t's monotonic clock
// It returns 0 for a missing reading.
func GetTimeMono(t time.Time) int64 {
	mono := timeMono(&t)
	if mono == 0 {
		return 0
	}
	return mono + startNano
}

// MonoSub returns the duration mono1-mono2.
func MonoSub(mono1, mono2 int64) time.Duration {
	return time.Duration(mono1 - mono2)
}
