package problem00034

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_searchRange(t *testing.T) {
	funs := map[string]func([]int, int) []int{
		"searchRange": searchRange,
	}

	tests := map[string]struct {
		nums   []int
		target int
		want   []int
	}{
		"normal1": {
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		"normal2": {
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		"normal3": {
			nums:   []int{},
			target: 0,
			want:   []int{-1, -1},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.nums, tt.target)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
