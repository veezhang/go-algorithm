package problem00905

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_sortArrayByParityII(t *testing.T) {
	funs := map[string]func([]int) []int{
		"sortArrayByParity1": sortArrayByParity1,
		"sortArrayByParity2": sortArrayByParity2,
	}

	tests := map[string]struct {
		A []int
	}{
		"normal1": {
			A: []int{3, 1, 2, 4},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(append([]int(nil), tt.A...))
				even := true
				for _, v := range got {
					if !even && 0 == v&1 {
						t.Errorf("%#v\n", got)
					}
					if even && 0 != v&1 {
						even = false
					}
				}
				ori := append([]int(nil), tt.A...)
				sort.Ints(ori)
				sort.Ints(got)
				diff := cmp.Diff(got, ori)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
