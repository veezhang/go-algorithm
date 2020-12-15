package problem00842

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_monotoneIncreasingDigits(t *testing.T) {
	funs := map[string]func(int) int{
		"monotoneIncreasingDigits1": monotoneIncreasingDigits1,
		"monotoneIncreasingDigits2": monotoneIncreasingDigits2,
		"monotoneIncreasingDigits3": monotoneIncreasingDigits3,
	}

	tests := map[string]struct {
		N    int
		want int
	}{
		"normal1": {
			N:    10,
			want: 9,
		},
		"normal2": {
			N:    1234,
			want: 1234,
		},
		"normal3": {
			N:    332,
			want: 299,
		},
		"normal4": {
			N:    999999999,
			want: 999999999,
		},
		"normal5": {
			N:    123456799,
			want: 123456799,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.N)
				want := tt.want

				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
