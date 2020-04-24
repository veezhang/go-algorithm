package problem00011

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_MaxArea(t *testing.T) {
	funs := map[string]func(height []int) int{
		"maxArea": maxArea,
	}

	tests := map[string]struct {
		height []int
		want   int
	}{
		"normal": {
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.height)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
