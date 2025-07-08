//go:build !go1.23 || allow_linkname

package exruntime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUptime(t *testing.T) {
	time.Sleep(1)
	assert.True(t, Uptime() > 0)
}
