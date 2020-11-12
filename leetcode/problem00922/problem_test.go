package problem00922

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_sortArrayByParityII(t *testing.T) {
	funs := map[string]func([]int) []int{
		"sortArrayByParityII1": sortArrayByParityII1,
		"sortArrayByParityII2": sortArrayByParityII2,
	}

	tests := map[string]struct {
		A []int
	}{
		"normal1": {
			A: []int{4, 2, 5, 7},
		},
		"normal2": {
			A: []int{2, 0, 3, 4, 1, 3},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(append([]int(nil), tt.A...))

				for i, v := range got {
					if i&1 != v&1 {
						t.Errorf("%d&1 != %d&1\n", i, v)
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
