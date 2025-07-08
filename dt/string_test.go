package dt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	cases := map[interface{}]string{
		1:         "1",
		true:      "true",
		'a':       "a",
		byte('a'): "a",
	}

	for v, s := range cases {
		t.Run("case", func(t *testing.T) {
			assert.Equal(t, s, ToString(v))
		})
	}
}
