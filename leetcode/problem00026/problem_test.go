package problem00026

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_RemoveDuplicates(t *testing.T) {
	funs := map[string]func(nums []int) int{
		"removeDuplicates": removeDuplicates,
	}

	tests := map[string]struct {
		nums []int
		want []int
	}{
		"normal": {
			nums: []int{1, 1, 2},
			want: []int{1, 2},
		},
		"normal1": {
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: []int{0, 1, 2, 3, 4},
		},
		"empty": {
			nums: []int{},
			want: []int{},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			// 深拷贝一份 ，因为会修改 head ，导致后面的数据出现问题
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				retn := fun(nums)
				got := nums[:retn]
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
