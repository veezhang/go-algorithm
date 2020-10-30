package problem00124

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_sumNumbers(t *testing.T) {
	funs := map[string]func(root *TreeNode) int{
		"maxPathSum": maxPathSum,
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
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: 6,
		},
		"normal2": {
			root: &TreeNode{
				Val: -10,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			want: 42,
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
