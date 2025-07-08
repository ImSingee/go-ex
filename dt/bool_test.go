package dt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBool(t *testing.T) {
	// All non-zero numbers are true
	trueCases := []interface{}{
		true,
		"1", "ok", "001",
		1, -1, 1.1,
	}
	falseCases := []interface{}{
		false,
		"0", "false", "000",
		0, 0.0, nil,
	}

	for _, c := range trueCases {
		t.Run(fmt.Sprintf("Should Be True: %#+v", c), func(t *testing.T) {
			b, ok := ParseBool(c)
			assert.True(t, ok)
			assert.True(t, b)
		})
	}

	for _, c := range falseCases {
		t.Run(fmt.Sprintf("Should Be False: %#+v", c), func(t *testing.T) {
			b, ok := ParseBool(c)
			assert.True(t, ok)
			assert.False(t, b)
		})
	}

}
