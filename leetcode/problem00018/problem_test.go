package problem00016

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// 已经排序号的 int Slice 的 Slice
type IntSortedSliceSlice [][]int

func (p IntSortedSliceSlice) Len() int {
	return len(p)
}
func (p IntSortedSliceSlice) Less(i, j int) bool {
	leni := len(p[i])
	lenj := len(p[j])

	for k := 0; k < leni && k < lenj; k++ {
		if p[i][k] != p[j][k] {
			return p[i][k] < p[j][k]
		}
	}

	return leni < lenj
}
func (p IntSortedSliceSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Test_FourSum(t *testing.T) {
	funs := map[string]func(nums []int, target int) [][]int{
		"fourSum": fourSum,
	}

	tests := map[string]struct {
		nums   []int
		target int
		want   [][]int
	}{
		"normal": {
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			want:   [][]int{{-1, 0, 0, 1}, {-2, -1, 1, 2}, {-2, 0, 0, 2}},
		},
		"repeat": {
			nums:   []int{-7, -2, 0, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 5, 5, 5, 5, 6, 8, 9, 9, 9, 9, 10, 10, 11},
			target: 2,
			want:   [][]int{{-7, -2, 0, 11}, {-7, -2, 2, 9}, {-7, -2, 3, 8}, {-7, -2, 5, 6}, {-7, 0, 3, 6}, {-7, 0, 4, 5}, {-7, 2, 2, 5}, {-7, 2, 3, 4}, {-7, 3, 3, 3}, {-2, 0, 2, 2}},
		},
		"aaa": {
			nums:   []int{0, 4, -5, 2, -2, 4, 2, -1, 4},
			target: 12,
			want:   [][]int{{0, 4, 4, 4}, {2, 2, 4, 4}},
		},
		"empty": {
			nums:   []int{},
			target: 0,
			want:   [][]int{},
		},
		"2num": {
			nums:   []int{-1, 1},
			target: 0,
			want:   [][]int{},
		},
		"3num": {
			nums:   []int{-1, 1, 2},
			target: 2,
			want:   [][]int{},
		},
		"max4<target": {
			nums:   []int{3, 4, 5, 6, 7, 8, 9, 10},
			target: 35,
			want:   [][]int{},
		},
		"max<target/4": {
			nums:   []int{3, 4, 5, 6, 7, 8, 9, 10},
			target: 17,
			want:   [][]int{},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.nums, tt.target)
				want := tt.want

				// 排序下，不然对比可能有问题
				for i := 0; i < len(got); i++ {
					sort.Ints(got[i])
				}
				sort.Sort(IntSortedSliceSlice(got))

				for i := 0; i < len(want); i++ {
					sort.Ints(want[i])
				}
				sort.Sort(IntSortedSliceSlice(want))

				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
