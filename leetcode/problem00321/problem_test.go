package problem00321

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_maxNumber(t *testing.T) {
	funs := map[string]func(nums1 []int, nums2 []int, k int) []int{
		"maxNumber": maxNumber,
	}

	tests := map[string]struct {
		nums1 []int
		nums2 []int
		k     int
		want  []int
	}{
		"normal1": {
			nums1: []int{3, 4, 6, 5},
			nums2: []int{9, 1, 2, 5, 8, 3},
			k:     5,
			want:  []int{9, 8, 6, 5, 3},
		},
		"normal2": {
			nums1: []int{6, 7},
			nums2: []int{6, 0, 4},
			k:     5,
			want:  []int{6, 7, 6, 0, 4},
		},
		"normal3": {
			nums1: []int{3, 9},
			nums2: []int{8, 9},
			k:     3,
			want:  []int{9, 8, 9},
		},
		"normal4": {
			nums1: []int{3, 4, 6, 5},
			nums2: []int{9, 1, 2, 5, 8, 3},
			k:     0,
			want:  nil,
		},
		"normal5": {
			nums1: []int{3, 4, 6, 5},
			nums2: []int{9, 1, 2, 5, 8, 3},
			k:     1,
			want:  []int{9},
		},
		"normal6": {
			nums1: []int{3, 4, 6, 5},
			nums2: []int{9, 1, 2, 5, 8, 3},
			k:     2,
			want:  []int{9, 8},
		},
		"normal7": {
			nums1: []int{3, 4, 6, 5},
			nums2: []int{9, 1, 2, 5, 8, 3},
			k:     8,
			want:  []int{9, 6, 5, 1, 2, 5, 8, 3},
		},
		"normal8": {
			nums1: []int{7, 6, 1, 9, 3, 2, 3, 1, 1},
			nums2: []int{4, 0, 9, 9, 0, 5, 5, 4, 7},
			k:     9,
			want:  []int{9, 9, 9, 7, 3, 2, 3, 1, 1},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.nums1, tt.nums2, tt.k)
				want := tt.want
				diff := cmp.Diff(got, want)
				fmt.Printf("%#v %#v\n", got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
