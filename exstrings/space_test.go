package exstrings

import (
	"testing"

	"github.com/ImSingee/tt"
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
			tt.AssertEqual(t, c[1], FoldSpace(c[0]))
		})
	}
}
