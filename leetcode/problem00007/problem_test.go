package problem00007

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Reverse(t *testing.T) {
	funs := map[string]func(x int) int{
		"reverse": reverse,
	}

	tests := map[string]struct {
		x    int
		want int
	}{
		"normal": {
			x:    123,
			want: 321,
		},
		"neg-normal": {
			x:    -123,
			want: -321,
		},
		"zero": {
			x:    0,
			want: 0,
		},
		"zero-end": {
			x:    102030,
			want: 30201,
		},
		// 2147483647
		"max-not-over": {
			x:    7463847412,
			want: 2147483647,
		},
		"max-over1": {
			x:    8463847412,
			want: 0,
		},
		"max-over2": {
			x:    1563847412,
			want: 0,
		},
		// -2147483648
		"min-not-over": {
			x:    -8463847412,
			want: -2147483648,
		},
		"min-over1": {
			x:    -9463847412,
			want: 0,
		},
		"min-over2": {
			x:    -1563847412,
			want: 0,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.x)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
