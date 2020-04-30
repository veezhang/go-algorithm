package problem00024

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mohae/deepcopy"
)

func Test_SwapPairs(t *testing.T) {
	funs := map[string]func(head *ListNode) *ListNode{
		"swapPairsRecursion": swapPairsRecursion,
		"swapPairsPointer":   swapPairsPointer,
	}

	tests := map[string]struct {
		head *ListNode
		want *ListNode
	}{
		"normal even": {
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
		},
		"normal odd": {
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
		"empty": {
			head: nil,
			want: nil,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			// 深拷贝一份 ，因为会修改 head ，导致后面的数据出现问题
			head := deepcopy.Copy(tt.head).(*ListNode)
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(head)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
