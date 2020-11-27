package problem00164

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_maximumGap(t *testing.T) {
	funs := map[string]func([]int, []int, []int, []int) int{
		"fourSumCount": fourSumCount,
	}

	tests := map[string]struct {
		A    []int
		B    []int
		C    []int
		D    []int
		want int
	}{
		"normal1": {
			A:    []int{1, 2},
			B:    []int{-2, -1},
			C:    []int{-1, 2},
			D:    []int{0, 2},
			want: 2,
		},
		"normal2": {
			A:    []int{1},
			B:    []int{-2},
			C:    []int{4},
			D:    []int{-3},
			want: 1,
		},
		"normal3": {
			A:    []int{1},
			B:    []int{-2},
			C:    []int{4},
			D:    []int{-5},
			want: 0,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.A, tt.B, tt.C, tt.D)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
