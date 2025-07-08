package dt

import (
	"github.com/ImSingee/tt"
	"testing"
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
			tt.AssertEqual(t, s, ToString(v))
		})
	}
}
