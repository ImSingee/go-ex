package extime

import (
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
