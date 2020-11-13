package problem00328

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_singleNumber(t *testing.T) {
	funs := map[string]func(head *ListNode) *ListNode{
		"oddEvenList1": oddEvenList1,
		"oddEvenList2": oddEvenList2,
	}

	tests := map[string]struct {
		head []int
		want []int
	}{
		"normal1": {
			head: []int{1, 2, 3, 4, 5},
			want: []int{1, 3, 5, 2, 4},
		},
		"normal2": {
			head: []int{2, 1, 3, 5, 6, 4, 7},
			want: []int{2, 3, 6, 7, 1, 5, 4},
		},
		"normal3": {
			head: []int{2, 1, 3, 5, 6, 4},
			want: []int{2, 3, 6, 1, 5, 4},
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
