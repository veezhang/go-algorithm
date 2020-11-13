package problem00725

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_singleNumber(t *testing.T) {
	funs := map[string]func(*ListNode, int) []*ListNode{
		"splitListToParts": splitListToParts,
	}

	tests := map[string]struct {
		root []int
		k    int
		want [][]int
	}{
		"normal1": {
			root: []int{1, 2, 3},
			k:    5,
			want: [][]int{{1}, {2}, {3}, {}, {}},
		},
		"normal2": {
			root: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			k:    3,
			want: [][]int{{1, 2, 3, 4}, {5, 6, 7}, {8, 9, 10}},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(NewListFromSlice(tt.root), tt.k)
				want := tt.want

				got1 := make([][]int, len(got))
				for i, r := range got {
					if nil != r {
						got1[i] = r.Slice()
					} else {
						got1[i] = []int{}
					}
				}

				diff := cmp.Diff(got1, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
