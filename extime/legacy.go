//go:build !go1.23 || allow_linkname

package extime

import (
	"time"

	_ "unsafe"
)

//go:linkname startNano time.startNano
var startNano int64

// StartMono returns the start monotonic time of the program
func StartMono() int64 {
	return startNano + 1
}

// GetTimeMono returns the mono part of t's monotonic clock
// It returns 0 for a missing reading.
func GetTimeMono(t time.Time) int64 {
	mono := timeMono(&t)
	if mono == 0 {
		return 0
	}
	return mono + startNano
}
