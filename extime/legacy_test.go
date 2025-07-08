//go:build !go1.23 || allow_linkname

package extime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStartNano(t *testing.T) {
	assert.NotEqual(t, int64(0), startNano)
}

func TestGetTimeMono(t *testing.T) {
	now := time.Now()

	assert.NotEqual(t, int64(0), GetTimeMono(now))

	time.Sleep(1)
	now2 := time.Now()
	assert.True(t, GetTimeMono(now2) > GetTimeMono(now))

	time.Sleep(1)
	nano3 := Monotonic()
	assert.True(t, nano3 > GetTimeMono(now2))
}
