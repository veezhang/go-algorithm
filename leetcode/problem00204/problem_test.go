package problem00204

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_countPrimes(t *testing.T) {
	funs := map[string]func(int) int{
		"countPrimes1": countPrimes1,
		"countPrimes2": countPrimes2,
	}

	tests := map[string]struct {
		n    int
		want int
	}{
		"normal1": {
			n:    0,
			want: 0,
		},
		"normal2": {
			n:    1,
			want: 0,
		},
		"normal3": {
			n:    2,
			want: 0,
		},
		"normal4": {
			n:    3,
			want: 1,
		},
		"normal5": {
			n:    10,
			want: 4,
		},
		"normal6": {
			n:    1500000,
			want: 114155,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.n)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
