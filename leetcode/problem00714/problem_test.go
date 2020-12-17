package problem00714

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_maxProfit(t *testing.T) {
	funs := map[string]func([]int, int) int{
		"maxProfit1": maxProfit1,
		"maxProfit2": maxProfit2,
		"maxProfit3": maxProfit3,
	}

	tests := map[string]struct {
		prices []int
		fee    int
		want   int
	}{
		"normal1": {
			prices: []int{1, 3, 2, 8, 4, 9},
			fee:    2,
			want:   8,
		},
		"normal2": {
			prices: []int{1, 3, 7, 5, 10, 3},
			fee:    3,
			want:   6,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.prices, tt.fee)
				want := tt.want

				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
