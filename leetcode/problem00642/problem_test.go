package problem00860

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_predictPartyVictory(t *testing.T) {
	funs := map[string]func(string) string{
		"predictPartyVictory": predictPartyVictory,
	}

	tests := map[string]struct {
		senate string
		want   string
	}{
		"normal1": {
			senate: "RD",
			want:   "Radiant",
		},
		"normal2": {
			senate: "RDD",
			want:   "Dire",
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.senate)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
