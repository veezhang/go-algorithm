package problem00017

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_LetterCombinations(t *testing.T) {
	funs := map[string]func(digits string) []string{
		"letterCombinationsDirect":    letterCombinationsDirect,
		"letterCombinationsBacktrack": letterCombinationsBacktrack,
	}

	tests := map[string]struct {
		digits string
		want   []string
	}{
		"normal": {
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		"2": {

			digits: "2",
			want:   []string{"a", "b", "c"},
		},
		"empty": {
			digits: "",
			want:   []string{},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.digits)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
