package problem00136

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_singleNumber(t *testing.T) {
	funs := map[string]func([]int) int{
		"singleNumber": singleNumber,
	}

	tests := map[string]struct {
		nums []int
		want int
	}{
		"normal1": {
			nums: []int{2, 2, 1},
			want: 1,
		},
		"normal2": {
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		"normal3": {
			nums: []int{5},
			want: 5,
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
