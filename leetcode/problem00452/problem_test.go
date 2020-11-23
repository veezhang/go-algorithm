package problem00452

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_insertionSortList(t *testing.T) {
	funs := map[string]func([][]int) int{
		"findMinArrowShots": findMinArrowShots,
	}

	tests := map[string]struct {
		points [][]int
		want   int
	}{
		"normal1": {
			points: [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
			want:   2,
		},
		"normal2": {
			points: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			want:   4,
		},
		"normal3": {
			points: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			want:   2,
		},
		"normal4": {
			points: [][]int{{1, 2}},
			want:   1,
		},
		"normal5": {
			points: [][]int{{2, 3}, {2, 3}},
			want:   1,
		},
		"normal6": {
			points: [][]int{},
			want:   0,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.points)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
