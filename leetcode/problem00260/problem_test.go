package problem00260

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_singleNumber(t *testing.T) {
	funs := map[string]func([]int) []int{
		"singleNumber": singleNumber,
	}

	tests := map[string]struct {
		nums []int
		want []int
	}{
		"normal1": {
			nums: []int{1, 2, 1, 3, 2, 5},
			want: []int{3, 5},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.nums)
				want := tt.want
				sort.Sort(sort.IntSlice(got))
				sort.Sort(sort.IntSlice(want))
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
