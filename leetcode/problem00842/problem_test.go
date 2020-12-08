package problem00842

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_splitIntoFibonacci(t *testing.T) {
	funs := map[string]func(S string) []int{
		"splitIntoFibonacci": splitIntoFibonacci,
	}

	tests := map[string]struct {
		S    string
		want []int
	}{
		"normal1": {
			S:    "123456579",
			want: []int{123, 456, 579},
		},
		"normal2": {
			S:    "11235813",
			want: []int{1, 1, 2, 3, 5, 8, 13},
		},
		"normal3": {
			S:    "112358130",
			want: []int{},
		},
		"normal4": {
			S:    "0123",
			want: []int{},
		},
		"normal5": {
			S:    "1101111",
			want: []int{11, 0, 11, 11},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.S)
				want := tt.want

				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
