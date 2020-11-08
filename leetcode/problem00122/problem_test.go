package problem00122

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_maxProfit(t *testing.T) {
	funs := map[string]func([]int) int{
		"maxProfit1": maxProfit1,
		"maxProfit2": maxProfit2,
		"maxProfit3": maxProfit3,
		"maxProfit4": maxProfit4,
		"maxProfit5": maxProfit5,
	}

	tests := map[string]struct {
		prices []int
		want   int
	}{
		"normal1": {
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   7,
		},
		"normal2": {
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		"normal4": {
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		"normal5": {
			prices: []int{2, 1, 2, 0, 1},
			want:   2,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.prices)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
