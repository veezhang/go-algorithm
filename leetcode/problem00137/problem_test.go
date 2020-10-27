package problem00137

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_singleNumber(t *testing.T) {
	funs := map[string]func([]int) int{
		"singleNumber1": singleNumber1,
		"singleNumber":  singleNumber,
	}

	tests := map[string]struct {
		nums []int
		want int
	}{
		"normal1": {
			nums: []int{2, 2, 3, 2},
			want: 3,
		},
		"normal2": {
			nums: []int{0, 1, 0, 1, 0, 1, 99},
			want: 99,
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
