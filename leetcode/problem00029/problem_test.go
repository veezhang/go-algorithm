package problem00029

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Divide(t *testing.T) {
	funs := map[string]func(dividend int, divisor int) (result int){
		"divide": divide,
	}

	tests := map[string]struct {
		dividend int
		divisor  int
		want     int
	}{
		"normal+": {
			dividend: 10,
			divisor:  3,
			want:     3,
		},
		"normal-": {
			dividend: 7,
			divisor:  -3,
			want:     -2,
		},
		"overflow": {
			dividend: math.MinInt32,
			divisor:  -1,
			want:     math.MaxInt32,
		},
		"divisor1": {
			dividend: 123456,
			divisor:  1,
			want:     123456,
		},
		"divisor-1": {
			dividend: 1234567,
			divisor:  -1,
			want:     -1234567,
		},
		"min": {
			dividend: math.MinInt32,
			divisor:  math.MinInt32 + 1,
			want:     1,
		},
		"max": {
			dividend: math.MaxInt32,
			divisor:  math.MaxInt32 - 1,
			want:     1,
		},
		"max-max": {
			dividend: math.MaxInt32,
			divisor:  math.MaxInt32,
			want:     1,
		},
		"min-min": {
			dividend: math.MaxInt32,
			divisor:  math.MaxInt32,
			want:     1,
		},
		"max-min": {
			dividend: math.MaxInt32,
			divisor:  math.MinInt32 + 1,
			want:     -1,
		},
		"min-man": {
			dividend: math.MinInt32 + 1,
			divisor:  math.MaxInt32,
			want:     -1,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.dividend, tt.divisor)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
