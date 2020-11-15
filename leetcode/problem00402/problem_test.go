package problem00402

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_singleNumber(t *testing.T) {
	funs := map[string]func(string, int) string{
		"removeKdigits": removeKdigits,
	}

	tests := map[string]struct {
		num  string
		k    int
		want string
	}{
		"normal1": {
			num:  "1432219",
			k:    3,
			want: "1219",
		},
		"normal2": {
			num:  "10200",
			k:    1,
			want: "200",
		},
		"normal3": {
			num:  "10",
			k:    2,
			want: "0",
		},
		"normal4": {
			num:  "10",
			k:    1,
			want: "0",
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.num, tt.k)
				want := tt.want

				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
