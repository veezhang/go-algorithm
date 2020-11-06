package problem01356

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_ladderLength(t *testing.T) {
	funs := map[string]func([]int) []int{
		"sortByBits1": sortByBits1,
		"sortByBits2": sortByBits2,
		"sortByBits3": sortByBits3,
		"sortByBits4": sortByBits4,
		"sortByBits5": sortByBits5,
		"sortByBits6": sortByBits6,
	}

	tests := map[string]struct {
		arr  []int
		want []int
	}{
		"normal1": {
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			want: []int{0, 1, 2, 4, 8, 3, 5, 6, 7},
		},
		"normal2": {
			arr:  []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1},
			want: []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
		},
		"normal3": {
			arr:  []int{10000, 10000},
			want: []int{10000, 10000},
		},
		"normal4": {
			arr:  []int{2, 3, 5, 7, 11, 13, 17, 19},
			want: []int{2, 3, 5, 17, 7, 11, 13, 19},
		},
		"normal5": {
			arr:  []int{10, 100, 1000, 10000},
			want: []int{10, 100, 10000, 1000},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.arr)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
