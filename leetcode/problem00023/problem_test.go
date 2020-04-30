package problem00023

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mohae/deepcopy"
)

func Test_MergeTwoLists(t *testing.T) {
	funs := map[string]func(lists []*ListNode) *ListNode{
		"mergeKListsFindMinInK":                  mergeKListsFindMinInK,
		"mergeKListsOneByOne":                    mergeKListsOneByOne,
		"mergeKListsDivideAndConquer":            mergeKListsDivideAndConquer,
		"mergeKListsDivideAndConquerNoRecursion": mergeKListsDivideAndConquerNoRecursion,
		"mergeKListsHeap":                        mergeKListsHeap,
	}

	tests := map[string]struct {
		lists []*ListNode
		want  *ListNode
	}{
		"normal": {
			lists: []*ListNode{
				{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
				{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
				{
					Val: 2,
					Next: &ListNode{
						Val: 6,
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 5,
										Next: &ListNode{
											Val: 6,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"empty": {
			lists: []*ListNode{
				nil,
				nil,
			},
			want: nil,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			// 深拷贝一份 ，因为会修改 lists ，导致后面的数据出现问题
			lists := deepcopy.Copy(tt.lists).([]*ListNode)
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(lists)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
