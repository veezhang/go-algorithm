package problem00001

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_TwoSum(t *testing.T) {
	funs := map[string]func(nums []int, target int) []int{
		"twoSum": twoSum,
	}

	tests := map[string]struct {
		nums   []int
		target int
		want   []int
	}{
		"normal": {
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		"not found": {
			nums:   []int{2, 7, 11, 15},
			target: 12,
			want:   nil,
		},
		"nil": {
			nums:   nil,
			target: 9,
			want:   nil,
		},
		"empty": {
			nums:   []int{},
			target: 9,
			want:   nil,
		},
		"one nums": {
			nums:   []int{1},
			target: 1,
			want:   nil,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.nums, tt.target)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
