package problem00015

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

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

func Test_ThreeSum(t *testing.T) {
	funs := map[string]func(nums []int) [][]int{
		"threeSumHash":       threeSumHash,
		"threeSumTwoPointer": threeSumTwoPointer,
	}

	tests := map[string]struct {
		nums []int
		want [][]int
	}{
		"normal": {
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, 0, 1}, {-1, -1, 2}},
		},
		"repeat": {
			nums: []int{-7, -2, 0, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 5, 5, 5, 5, 6, 8, 9, 9, 9, 9, 10, 10, 11},
			want: [][]int{{-7, -2, 9}, {-7, 2, 5}, {-7, 3, 4}, {-2, 0, 2}},
		},
		"empty": {
			nums: []int{},
			want: [][]int{},
		},
		"2num": {
			nums: []int{-1, 1},
			want: [][]int{},
		},
		"all<0": {
			nums: []int{-1, -1, -1},
			want: [][]int{},
		},
		"all>0": {
			nums: []int{1, 1, 1},
			want: [][]int{},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.nums)
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

// go test -bench=. -count 10 -cpuprofile=profile.out ./leetcode/problem00015/...
// go tool pprof -no_browser -http=0.0.0.0:8080 profile.out
// runtime.mapaccess2_fast64 和 runtime.mapassgin_fast64 占用大量的时间

func Benchmark_threeSumHash(b *testing.B) {
	rand.Seed(time.Now().Unix())
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nums[i] = rand.Intn(10000) - 5000
	}

	for i := 0; i < b.N; i++ {
		threeSumHash(nums)
	}
}
