package dt

import (
	"fmt"
	"github.com/ImSingee/tt"
	"testing"
)

func TestInSlice(t *testing.T) {
	trueCases := [][2]interface{}{
		{[]interface{}{nil}, nil},
		{[]string{"a", "b", "c"}, "a"},
		{[]int{1, 3, 5}, 1},
		{[]interface{}{1, 3, 5, "7"}, "7"},
	}

	for i, c := range trueCases {
		t.Run(fmt.Sprintf("true - %d", i), func(t *testing.T) {
			tt.AssertTrue(t, InSlice(c[0], c[1]))
		})
	}

	falseCases := [][2]interface{}{
		{[]interface{}{}, nil},
		{[]string{"a", "b", "c"}, "e"},
		{[]int{1, 3, 5}, 7},
		{[]int{1, 3, 5}, "1"},
	}

	for i, c := range falseCases {
		t.Run(fmt.Sprintf("false - %d", i), func(t *testing.T) {
			tt.AssertFalse(t, InSlice(c[0], c[1]))
		})
	}
}
