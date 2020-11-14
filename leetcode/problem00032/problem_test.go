package problem00032

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_maxProfit(t *testing.T) {
	funs := map[string]func(s string) (ans int){
		"longestValidParentheses1": longestValidParentheses1,
		"longestValidParentheses2": longestValidParentheses2,
		"longestValidParentheses3": longestValidParentheses3,
	}

	tests := map[string]struct {
		s    string
		want int
	}{
		"normal1": {
			s:    "(()",
			want: 2,
		},
		"normal2": {
			s:    ")()())",
			want: 4,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.s)
				want := tt.want
				if diff := cmp.Diff(got, want); diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
