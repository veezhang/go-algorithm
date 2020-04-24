package problem00007

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_IsPalindrome(t *testing.T) {
	funs := map[string]func(x int) bool{
		"isPalindromeCompareHightAndLow": isPalindromeCompareHightAndLow,
		"isPalindromeReverseHalf":        isPalindromeReverseHalf,
	}

	tests := map[string]struct {
		x    int
		want bool
	}{
		"normal": {
			x:    12321,
			want: true,
		},
		"neg-normal": {
			x:    -12321,
			want: false,
		},
		"zero": {
			x:    0,
			want: true,
		},
		"not": {
			x:    123210,
			want: false,
		},
		"1001": {
			x:    1001,
			want: true,
		},
		"1": {
			x:    1,
			want: true,
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
