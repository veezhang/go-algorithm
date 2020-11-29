package problem00976

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_maximumGap(t *testing.T) {
	funs := map[string]func([]int) int{
		"largestPerimeter": largestPerimeter,
	}

	tests := map[string]struct {
		A    []int
		want int
	}{
		"normal1": {
			A:    []int{2, 1, 2},
			want: 5,
		},
		"normal2": {
			A:    []int{1, 2, 1},
			want: 0,
		},
		"normal3": {
			A:    []int{3, 2, 3, 4},
			want: 10,
		},
		"normal4": {
			A:    []int{3, 6, 2, 3},
			want: 8,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.A)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
