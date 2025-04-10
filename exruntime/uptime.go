//go:build !go1.23

package exruntime

import (
	"github.com/ImSingee/go-ex/extime"
	"time"
)

// Uptime returns the program uptime.
func Uptime() time.Duration {
	return time.Duration(extime.Monotonic() - extime.StartMono())
}
