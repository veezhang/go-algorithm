package problem00031

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_maxProfit(t *testing.T) {
	funs := map[string]func(nums []int){
		"nextPermutation": nextPermutation,
	}

	tests := map[string]struct {
		nums []int
		want []int
	}{
		"normal1": {
			nums: []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		"normal2": {
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		"normal3": {
			nums: []int{1, 1, 5},
			want: []int{1, 5, 1},
		},
		"normal4": {
			nums: []int{1},
			want: []int{1},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				fun(tt.nums)
				got := tt.nums
				want := tt.want
				if diff := cmp.Diff(got, want); diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
