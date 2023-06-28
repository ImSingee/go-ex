package exruntime

import (
	"github.com/ImSingee/tt"
	"testing"
)

func TestUptime(t *testing.T) {
	tt.AssertTrue(t, Uptime() > 0)
}
