package dt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
			assert.True(t, InSlice(c[0], c[1]))
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
			assert.False(t, InSlice(c[0], c[1]))
		})
	}
}
