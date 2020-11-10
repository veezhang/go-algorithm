package problem00973

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_kClosest(t *testing.T) {
	funs := map[string]func([][]int, int) [][]int{
		"kClosest1": kClosest1,
		"kClosest2": kClosest2,
		"kClosest3": kClosest3,
		"kClosest4": kClosest4,
		"kClosest5": kClosest5,
	}

	tests := map[string]struct {
		points [][]int
		K      int
		want   [][]int
	}{
		"normal1": {
			points: [][]int{{1, 3}, {-2, 2}},
			K:      1,
			want:   [][]int{{-2, 2}},
		},
		"normal2": {
			points: [][]int{{1, 3}, {-2, 2}},
			K:      2,
			want:   [][]int{{1, 3}, {-2, 2}},
		},
		"normal3": {
			points: [][]int{{1, 3}, {-2, 2}},
			K:      3,
			want:   [][]int{{1, 3}, {-2, 2}},
		},
		"normal4": {
			points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
			K:      2,
			want:   [][]int{{3, 3}, {-2, 4}},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.points, tt.K)
				want := tt.want
				sort.Slice(want, func(i, j int) bool {
					pi, pj := want[i], want[j]
					return pi[0]*pi[0] < pj[0]*pj[0]
				})
				sort.Slice(got, func(i, j int) bool {
					pi, pj := got[i], got[j]
					return pi[0]*pi[0] < pj[0]*pj[0]
				})
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
