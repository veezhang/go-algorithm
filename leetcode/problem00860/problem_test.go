package problem00860

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_lemonadeChange(t *testing.T) {
	funs := map[string]func([]int) bool{
		"lemonadeChange": lemonadeChange,
	}

	tests := map[string]struct {
		bills []int
		want  bool
	}{
		"normal1": {
			bills: []int{5, 5, 5, 10, 20},
			want:  true,
		},
		"normal2": {
			bills: []int{5, 5, 10},
			want:  true,
		},
		"normal3": {
			bills: []int{10, 10},
			want:  false,
		},
		"normal4": {
			bills: []int{5, 5, 10, 10, 20},
			want:  false,
		},
		"normal5": {
			bills: []int{},
			want:  true,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.bills)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
