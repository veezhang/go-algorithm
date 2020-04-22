package problem00004

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_FindMedianSortedArrays(t *testing.T) {
	funs := map[string]func(nums1 []int, nums2 []int) float64{
		"findMedianSortedArraysKth":    findMedianSortedArraysKth,
		"findMedianSortedArraysBinary": findMedianSortedArraysBinary,
	}

	tests := map[string]struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		"odd1": {
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.0,
		},
		"odd2": {
			nums1: []int{2},
			nums2: []int{1, 3},
			want:  2.0,
		},
		"even1": {
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
		"even2": {
			nums1: []int{3, 4},
			nums2: []int{1, 2},
			want:  2.5,
		},
		"nums1 one": {
			nums1: []int{3},
			nums2: []int{1, 2, 4, 5, 6},
			want:  3.5,
		},
		"nums2 one": {
			nums1: []int{1, 2, 4, 5, 6},
			nums2: []int{3},
			want:  3.5,
		},
		"nums1 one+": {
			nums1: []int{3},
			nums2: []int{1, 2, 4, 5, 6, 7, 8, 9, 10},
			want:  5.5,
		},
		"nums2 one+": {
			nums1: []int{1, 2, 4, 5, 6, 7, 8, 9, 10},
			nums2: []int{3},
			want:  5.5,
		},
		"m1in1 m2in1": {
			nums1: []int{1, 2, 3, 4},
			nums2: []int{5, 6},
			want:  3.5,
		},
		"m1in2 m2in2": {
			nums1: []int{5, 6},
			nums2: []int{1, 2, 3, 4},
			want:  3.5,
		},
		"only nums1 odd": {
			nums1: []int{1, 2, 3, 4, 5},
			nums2: []int{},
			want:  3,
		},
		"only nums2 odd": {
			nums1: []int{},
			nums2: []int{1, 2, 3, 4, 5},
			want:  3,
		},
		"only nums1 even": {
			nums1: []int{1, 2, 3, 4},
			nums2: []int{},
			want:  2.5,
		},
		"only nums2 even": {
			nums1: []int{},
			nums2: []int{1, 2, 3, 4},
			want:  2.5,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.nums1, tt.nums2)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
