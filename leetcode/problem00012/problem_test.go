package problem00012

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_IntToRoman(t *testing.T) {
	funs := map[string]func(num int) string{
		"intToRomanStaticTable":      intToRomanStaticTable,
		"intToRomanGreedyAlgorithms": intToRomanGreedyAlgorithms,
	}

	tests := map[string]struct {
		num  int
		want string
	}{
		"3": {
			num:  3,
			want: "III",
		},
		"4": {
			num:  4,
			want: "IV",
		},
		"9": {
			num:  9,
			want: "IX",
		},
		"58": {
			num:  58,
			want: "LVIII",
		},
		"1994": {
			num:  1994,
			want: "MCMXCIV",
		},
		"2003": {
			num:  2003,
			want: "MMIII",
		},
		"3456": {
			num:  3456,
			want: "MMMCDLVI",
		},
		"3987": {
			num:  3987,
			want: "MMMCMLXXXVII",
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.num)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
