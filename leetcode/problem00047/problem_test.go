package problem00047

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_maxProfit(t *testing.T) {
	funs := map[string]func(nums []int) [][]int{
		"permuteUnique": permuteUnique,
	}

	tests := map[string]struct {
		nums []int
		want [][]int
	}{
		"normal1": {
			nums: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		"normal2": {
			nums: []int{1, 1, 2},
			want: [][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.nums)
				want := tt.want
				less := func(data [][]int, i, j int) bool {
					xi, xj := data[i], data[j]
					for k := 0; k < len(xi); k++ {
						if xi[k] < xj[k] {
							return true
						}
						if xi[k] > xj[k] {
							return false
						}
					}
					return false
				}
				sort.Slice(got, func(i, j int) bool {
					return less(got, i, j)
				})
				sort.Slice(want, func(i, j int) bool {
					return less(want, i, j)
				})
				if diff := cmp.Diff(got, want); diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
