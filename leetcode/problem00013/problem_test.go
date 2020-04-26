package problem00013

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_IntToRoman(t *testing.T) {
	funs := map[string]func(s string) int{
		"romanToInt": romanToInt,
	}

	tests := map[string]struct {
		s    string
		want int
	}{
		"3": {
			s:    "III",
			want: 3,
		},
		"4": {
			s:    "IV",
			want: 4,
		},
		"9": {
			s:    "IX",
			want: 9,
		},
		"58": {
			s:    "LVIII",
			want: 58,
		},
		"1994": {
			s:    "MCMXCIV",
			want: 1994,
		},
		"2003": {
			s:    "MMIII",
			want: 2003,
		},
		"3456": {
			s:    "MMMCDLVI",
			want: 3456,
		},
		"3987": {
			s:    "MMMCMLXXXVII",
			want: 3987,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.s)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
