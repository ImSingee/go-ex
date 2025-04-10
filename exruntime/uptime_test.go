//go:build !go1.23 || allow_linkname

package exruntime

import (
	"github.com/ImSingee/tt"
	"testing"
	"time"
)

func TestUptime(t *testing.T) {
	time.Sleep(1)
	tt.AssertTrue(t, Uptime() > 0)
}
