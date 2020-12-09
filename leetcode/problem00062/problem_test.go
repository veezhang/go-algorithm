package problem00062

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_uniquePaths(t *testing.T) {
	funs := map[string]func(m int, n int) int{
		// "uniquePaths1": uniquePaths1,
		"uniquePaths2": uniquePaths2,
	}

	tests := map[string]struct {
		m    int
		n    int
		want int
	}{
		"normal1": {
			m:    3,
			n:    7,
			want: 28,
		},
		"normal2": {
			m:    3,
			n:    2,
			want: 3,
		},
		"normal3": {
			m:    7,
			n:    3,
			want: 28,
		},
		"normal4": {
			m:    3,
			n:    3,
			want: 6,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.m, tt.n)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
