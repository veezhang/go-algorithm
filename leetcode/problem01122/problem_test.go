package problem01122

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_singleNumber(t *testing.T) {
	funs := map[string]func([]int, []int) []int{
		"relativeSortArray1": relativeSortArray1,
		"relativeSortArray2": relativeSortArray2,
	}

	tests := map[string]struct {
		arr1 []int
		arr2 []int
		want []int
	}{
		"normal1": {
			arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			arr2: []int{2, 1, 4, 3, 9, 6},
			want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		"normal2": {
			arr1: []int{33, 22, 48, 4, 39, 36, 41, 47, 15, 45},
			arr2: []int{22, 33, 48, 4},
			want: []int{22, 33, 48, 4, 15, 36, 39, 41, 45, 47},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.arr1, tt.arr2)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
