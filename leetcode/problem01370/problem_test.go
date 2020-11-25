package problem00222

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_sortString(t *testing.T) {
	funs := map[string]func(string) string{
		"sortString": sortString,
	}

	tests := map[string]struct {
		s    string
		want string
	}{
		"normal1": {
			s:    "aaaabbbbcccc",
			want: "abccbaabccba",
		},
		"normal2": {
			s:    "rat",
			want: "art",
		},
		"normal3": {
			s:    "leetcode",
			want: "cdelotee",
		},
		"normal4": {
			s:    "ggggggg",
			want: "ggggggg",
		},
		"normal5": {
			s:    "spo",
			want: "ops",
		},
		"normal6": {
			s:    "abc",
			want: "abc",
		},
		"normal7": {
			s:    "a",
			want: "a",
		},
		"normal8": {
			s:    "",
			want: "",
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.s)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
