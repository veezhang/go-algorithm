package problem00016

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_ThreeSumClosest(t *testing.T) {
	funs := map[string]func(nums []int, target int) int{
		"threeSumClosest": threeSumClosest,
	}

	tests := map[string]struct {
		nums   []int
		target int
		want   int
	}{
		"normal": {
			nums:   []int{-1, 2, 1, -4},
			target: 1,
			want:   2,
		},
		"normal1": {
			nums:   []int{1, 1, 1, 0},
			target: 100,
			want:   3,
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
