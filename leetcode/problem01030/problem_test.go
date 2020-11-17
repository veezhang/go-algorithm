package problem00406

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_allCellsDistOrder(t *testing.T) {
	funs := map[string]func(R int, C int, r0 int, c0 int) [][]int{
		"allCellsDistOrder1": allCellsDistOrder1,
		"allCellsDistOrder2": allCellsDistOrder2,
		"allCellsDistOrder3": allCellsDistOrder3,
		"allCellsDistOrder4": allCellsDistOrder4,
	}

	tests := map[string]struct {
		R    int
		C    int
		r0   int
		c0   int
		want [][]int
	}{
		"normal1": {
			R:    1,
			C:    2,
			r0:   0,
			c0:   0,
			want: [][]int{{0, 0}, {0, 1}},
		},
		"normal2": {
			R:    2,
			C:    2,
			r0:   0,
			c0:   1,
			want: [][]int{{0, 1}, {0, 0}, {1, 1}, {1, 0}},
		},
		"normal3": {
			R:    2,
			C:    3,
			r0:   1,
			c0:   2,
			want: [][]int{{1, 2}, {0, 2}, {1, 1}, {0, 1}, {1, 0}, {0, 0}},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.R, tt.C, tt.r0, tt.c0)
				want := tt.want

				abs := func(x int) int {
					if x < 0 {
						return -x
					}
					return x
				}

				// 先比较是否是升序
				sort.SliceIsSorted(got, func(i, j int) bool {
					x, y := got[i], got[j]
					lx, ly := abs(x[0]-tt.r0)+abs(x[0]-tt.c0), abs(y[0]-tt.r0)+abs(y[0]-tt.c0)
					return lx < ly
				})

				// 在将元素排序，看看是否元素一样
				less := func(i, j int) bool {
					x, y := got[i], got[j]
					lx, ly := abs(x[0]-tt.r0)+abs(x[1]-tt.c0), abs(y[0]-tt.r0)+abs(y[1]-tt.c0)
					// 先比较距离，在比较 r ， 在比较 c
					return lx < ly || (lx == ly && x[0] < y[0]) || (lx == ly && x[0] == y[0] && x[1] < y[1])
				}
				sort.Slice(got, less)
				sort.Slice(want, less)

				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
