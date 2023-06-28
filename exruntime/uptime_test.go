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
