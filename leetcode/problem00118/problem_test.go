package problem00118

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_generate(t *testing.T) {
	funs := map[string]func(int) [][]int{
		"generate": generate,
	}

	tests := map[string]struct {
		numRows int
		want    [][]int
	}{
		"normal1": {
			numRows: 5,
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.numRows)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
