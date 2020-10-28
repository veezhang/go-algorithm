package problem00287

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_singleNumber(t *testing.T) {
	funs := map[string]func([]int) int{
		"findDuplicate1": findDuplicate1,
		"findDuplicate2": findDuplicate2,
		"findDuplicate3": findDuplicate3,
	}

	tests := map[string]struct {
		nums []int
		want int
	}{
		"normal1": {
			nums: []int{1, 3, 4, 2, 2},
			want: 2,
		},
		"normal2": {
			nums: []int{3, 1, 3, 4, 2},
			want: 3,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.nums)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
