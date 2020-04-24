package problem00010

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_IsMatch(t *testing.T) {
	funs := map[string]func(s string, p string) bool{
		"isMatchBackTrace": isMatchBackTrace,
		"isMatchDP":        isMatchDP,
		"isMatchDP1":       isMatchDP1,
	}

	tests := map[string]struct {
		s    string
		p    string
		want bool
	}{
		"aa-a": {
			s:    "aa",
			p:    "a",
			want: false,
		},
		"aa-a*": {
			s:    "aa",
			p:    "a*",
			want: true,
		},
		"ab-.*": {
			s:    "aa",
			p:    ".*",
			want: true,
		},
		"aab-c*a*b": {
			s:    "aab",
			p:    "c*a*b",
			want: true,
		},
		"mississippi-mis*is*p*.": {
			s:    "mississippi",
			p:    "mis*is*p*.",
			want: false,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.s, tt.p)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
