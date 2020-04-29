package problem00022

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_GenerateParenthesis(t *testing.T) {
	funs := map[string]func(n int) []string{
		"generateParenthesis": generateParenthesis,
	}

	tests := map[string]struct {
		n    int
		want []string
	}{
		"normal": {
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		"zore": {
			n:    0,
			want: []string{},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.n)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
