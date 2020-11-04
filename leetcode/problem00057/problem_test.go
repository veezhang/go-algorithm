package problem00057

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_insert(t *testing.T) {
	funs := map[string]func(intervals [][]int, newInterval []int) [][]int{
		"insert": insert,
	}

	tests := map[string]struct {
		sintervals  [][]int
		newInterval []int
		want        [][]int
	}{
		"normal1": {
			sintervals:  [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			want:        [][]int{{1, 5}, {6, 9}},
		},
		"normal2": {
			sintervals:  [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			want:        [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.sintervals, tt.newInterval)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
