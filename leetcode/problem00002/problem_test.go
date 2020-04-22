package problem00002

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func newListNodeFromSlice(s []int) *ListNode {

	head := &ListNode{}
	temp := head
	for i := 0; i < len(s); i++ {
		temp.Next = &ListNode{
			Val: s[i],
		}
		temp = temp.Next
	}

	return head.Next
}

func Test_AddTwoNumbers(t *testing.T) {
	funs := map[string]func(l1 *ListNode, l2 *ListNode) *ListNode{
		"addTwoNumbers": addTwoNumbers,
	}

	tests := map[string]struct {
		l1_s   []int
		l2_s   []int
		want_s []int
	}{
		"normal": {
			l1_s:   []int{2, 4, 3},
			l2_s:   []int{5, 6, 4},
			want_s: []int{7, 0, 8},
		},
		"999999+1": {
			l1_s:   []int{9, 9, 9, 9, 9, 9},
			l2_s:   []int{1},
			want_s: []int{0, 0, 0, 0, 0, 0, 1},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				l1 := newListNodeFromSlice(tt.l1_s)
				l2 := newListNodeFromSlice(tt.l2_s)
				got := fun(l1, l2)
				want := newListNodeFromSlice(tt.want_s)
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
