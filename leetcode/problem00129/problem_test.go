package problem00129

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_sumNumbers(t *testing.T) {
	funs := map[string]func(root *TreeNode) int{
		"sumNumbersDFS": sumNumbersDFS,
		"sumNumbersBFS": sumNumbersBFS,
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
			want: 25,
		},
		"normal2": {
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 0,
				},
			},
			want: 1026,
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
