package exstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoldSpace(t *testing.T) {
	cases := [][2]string{
		{"abc", "abc"},
		{"a  b c", "a b c"},
		{"   a  b c   ", " a b c "},
		{"   a  b c", " a b c"},
		{"a  b c   ", "a b c "},
	}

	for _, c := range cases {
		t.Run(c[0], func(t *testing.T) {
			assert.Equal(t, c[1], FoldSpace(c[0]))
		})
	}
}
