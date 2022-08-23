package exstrings

import (
	"github.com/ImSingee/tt"
	"testing"
)

func TestLeftPad(t *testing.T) {
	cases := []struct {
		name      string
		s         string
		pad       string
		minLength int
		want      string
	}{
		{
			name:      `normal`,
			s:         "abc",
			pad:       "x",
			minLength: 5,
			want:      "xxabc",
		},
		{
			name:      `equal`,
			s:         "abc",
			pad:       "x",
			minLength: 3,
			want:      "abc",
		},
		{
			name:      `less`,
			s:         "abc",
			pad:       "x",
			minLength: 2,
			want:      "abc",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := LeftPad(c.s, c.pad, c.minLength)

			tt.AssertEqual(t, got, c.want)
		})
	}
}
