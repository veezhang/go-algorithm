package problem00514

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_kClosest(t *testing.T) {
	funs := map[string]func(ring string, key string) int{
		"findRotateSteps1": findRotateSteps1,
		"findRotateSteps2": findRotateSteps2,
	}

	tests := map[string]struct {
		ring string
		key  string
		want int
	}{
		"normal1": {
			ring: "godding",
			key:  "gd",
			want: 4,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.ring, tt.key)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
