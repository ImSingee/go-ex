package exstrings

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListToInt(t *testing.T) {
	cases := []struct {
		name      string
		in        []string
		want      []int
		willError bool
	}{
		{
			name:      `empty`,
			in:        nil,
			want:      []int{},
			willError: false,
		},
		{
			name: `one`,
			in:   []string{"16"},
			want: []int{16},
		},
		{
			name: `contains space`,
			in:   []string{" 16", "17 ", " 18 "},
			want: []int{16, 17, 18},
		},
		{
			name:      `invalid-1`,
			in:        []string{"hh"},
			want:      nil,
			willError: true,
		},
		{
			name:      `invalid-2`,
			in:        []string{"2hh"},
			want:      nil,
			willError: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := ListToInt(c.in)

			if !c.willError {
				assert.Nil(t, err)

				assert.Equal(t, c.want, got)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	base := []string{"h", "e", "l"}
	cmp := []string{"k", "l", "p"}

	add, sub, equal := Diff(cmp, base)
	sort.Strings(add)
	sort.Strings(sub)
	sort.Strings(equal)
	assert.Equal(t, []string{"k", "p"}, add)
	assert.Equal(t, []string{"e", "h"}, sub)
	assert.Equal(t, []string{"l"}, equal)
}
