package problem00008

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_MyAtoi(t *testing.T) {
	funs := map[string]func(str string) int{
		"myAtoiNormal": myAtoiNormal,
		"myAtoiDFA":    myAtoiDFA,
	}

	tests := map[string]struct {
		str  string
		want int
	}{
		"normal": {
			str:  "42",
			want: 42,
		},
		"neg-normal": {
			str:  "    -42",
			want: -42,
		},
		"empty": {
			str:  "",
			want: 0,
		},
		// 2147483647
		"max-not-over": {
			str:  "2147483647",
			want: 2147483647,
		},
		"max-over1": {
			str:  "2147483648",
			want: 2147483647,
		},
		"max-over2": {
			str:  "2147483657",
			want: 2147483647,
		},
		// -2147483648
		"min-not-over": {
			str:  "-2147483648",
			want: -2147483648,
		},
		"min-over1": {
			str:  "-2147483649",
			want: -2147483648,
		},
		"min-over2": {
			str:  "-2147483658",
			want: -2147483648,
		},
		"words and 987": {
			str:  "words and 987",
			want: 0,
		},
		"987 and words": {
			str:  "987 and words",
			want: 987,
		},
		"-987 7 ": {
			str:  "-987 7 ",
			want: -987,
		},
		" - 987 7 ": {
			str:  " - 987 7 ",
			want: 0,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.str)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
