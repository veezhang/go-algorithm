package problem00147

import (
	"fmt"
	"testing"

	. "github/veezhang/go-algorithm/leetcode/common/list"

	"github.com/google/go-cmp/cmp"
)

func Test_sortList(t *testing.T) {
	funs := map[string]func(head *ListNode) *ListNode{
		"sortList": sortList,
	}

	tests := map[string]struct {
		head []int
		want []int
	}{
		"normal1": {
			head: []int{4, 2, 1, 3},
			want: []int{1, 2, 3, 4},
		},
		"normal2": {
			head: []int{1, 2, 3, 4},
			want: []int{1, 2, 3, 4},
		},
		"normal3": {
			head: []int{4, 3, 2, 1},
			want: []int{1, 2, 3, 4},
		},
		"normal4": {
			head: []int{-1, 5, 3, 4, 0},
			want: []int{-1, 0, 3, 4, 5},
		},
		"normal5": {
			head: []int{1},
			want: []int{1},
		},
		"normal6": {
			head: []int{},
			want: []int{},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(NewListFromSlice(tt.head))
				want := tt.want
				diff := cmp.Diff(got.Slice(), want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
