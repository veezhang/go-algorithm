package problem00025

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_ListNode_ToSlice(t *testing.T) {
	tests := map[string]struct {
		s []int
		l *ListNode
	}{
		"normal": {
			s: []int{1, 2, 3, 4, 5},
			l: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
		"empty": {
			s: []int{},
			l: nil,
		},
		"siganl node": {
			s: []int{1},
			l: &ListNode{
				Val: 1,
			},
		},
	}

	// 测试ListNode.ToSlice
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.l.ToSlice()
			want := tt.s
			diff := cmp.Diff(got, want)
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func Test_ListNode_NewFromSlice(t *testing.T) {
	tests := map[string]struct {
		s []int
		l *ListNode
	}{
		"normal": {
			s: []int{1, 2, 3, 4, 5},
			l: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
		"empty": {
			s: []int{},
			l: nil,
		},
		"siganl node": {
			s: []int{1},
			l: &ListNode{
				Val: 1,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := NewFromSlice(tt.s)
			want := tt.l
			diff := cmp.Diff(got, want)
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func Test_ListNode_Equal(t *testing.T) {
	tests := map[string]struct {
		s  []int
		l  *ListNode
		eq bool
	}{
		"normal true": {
			s: []int{1, 2, 3, 4, 5},
			l: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			eq: true,
		},
		"normal false": {
			s: []int{1, 2, 3, 4, 5},
			l: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 4,
							},
						},
					},
				},
			},
			eq: false,
		},
		"empty true": {
			s:  []int{},
			l:  nil,
			eq: true,
		},
		"empty false": {
			s:  []int{},
			l:  &ListNode{Val: 0},
			eq: false,
		},
		"length diff": {
			s: []int{1, 2, 3, 4, 5, 6},
			l: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			eq: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			l := NewFromSlice(tt.s)
			got := l.Equal(tt.l)
			want := tt.eq
			diff := cmp.Diff(got, want)
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func Test_ReverseList(t *testing.T) {
	tests := []struct {
		s  []int
		rs []int
	}{
		{
			s:  []int{1},
			rs: []int{1},
		},
		{
			s:  []int{1, 2, 3, 4, 5},
			rs: []int{5, 4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run("ListNode.ReverseList", func(t *testing.T) {
			l := NewFromSlice(tt.s)
			got := reverseList(l)
			want := NewFromSlice(tt.rs)

			if !got.Equal(want) {
				t.Errorf("ListNode.ReverseList got %v, want %v", got.ToSlice(), want.ToSlice())
			}
		})
	}
}

func Test_ReverseKGroup(t *testing.T) {
	tests := map[string]struct {
		s   []int
		k   int
		rks []int
	}{
		"normal k1": {
			s:   []int{1, 2, 3, 4, 5},
			k:   1,
			rks: []int{1, 2, 3, 4, 5},
		},
		"normal k2": {
			s:   []int{1, 2, 3, 4, 5},
			k:   2,
			rks: []int{2, 1, 4, 3, 5},
		},
		"normal k3": {
			s:   []int{1, 2, 3, 4, 5},
			k:   3,
			rks: []int{3, 2, 1, 4, 5},
		},
		"not left": {
			s:   []int{1, 2, 3, 4, 5, 6},
			k:   3,
			rks: []int{3, 2, 1, 6, 5, 4},
		},
		"empty k1": {
			s:   []int{},
			k:   3,
			rks: []int{},
		},
		"empty k2": {
			s:   []int{},
			k:   3,
			rks: []int{},
		},
		"empty k3": {
			s:   []int{},
			k:   3,
			rks: []int{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			l := NewFromSlice(tt.s)
			got := reverseKGroup(l, tt.k)
			want := NewFromSlice(tt.rks)
			diff := cmp.Diff(got, want)
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
