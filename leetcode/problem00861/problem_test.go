package problem00861

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_matrixScore(t *testing.T) {
	funs := map[string]func(A [][]int) (ans int){
		"matrixScore": matrixScore,
	}

	tests := map[string]struct {
		A    [][]int
		want int
	}{
		"normal1": {
			A: [][]int{
				{0, 0, 1, 1},
				{1, 0, 1, 0},
				{1, 1, 0, 0},
			},
			want: 39,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.A)
				want := tt.want

				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
