package problem00222

import (
	"fmt"
	"testing"

	. "github/veezhang/go-algorithm/leetcode/common/tree"

	"github.com/google/go-cmp/cmp"
)

func Test_countNodes(t *testing.T) {
	funs := map[string]func(*TreeNode) int{
		"countNodes1": countNodes1,
		"countNodes2": countNodes2,
		"countNodes3": countNodes3,
	}

	tests := map[string]struct {
		root *TreeNode
		want int
	}{
		"normal1": {
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
				},
			},
			want: 6,
		},
		"normal2": {
			root: nil,
			want: 0,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.root)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
