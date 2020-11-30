package problem00767

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_reorganizeString(t *testing.T) {
	funs := map[string]func(string) string{
		"reorganizeString": reorganizeString,
	}

	tests := map[string]struct {
		S    string
		want string
	}{
		"normal1": {
			S:    "aab",
			want: "aba",
		},
		"normal2": {
			S:    "aaab",
			want: "",
		},
		"normal3": {
			S:    "aaaaaabbbbbccc",
			want: "abababcabacabc",
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.S)
				want := tt.want

				if len(got) != len(want) {
					t.Errorf("Length diff (%s) %d != (%s) %d", got, len(got), want, len(want))
				}
				var cnt1, cnt2 [26]int

				for _, ch := range []byte(got) {
					cnt1[ch-'a']++
				}

				for _, ch := range []byte(want) {
					cnt2[ch-'a']++
				}

				diff := cmp.Diff(cnt1, cnt2)
				if diff != "" {
					t.Errorf(diff)
				}

				for i := 1; i < len(got); i++ {
					if got[i-1] == got[i] {
						t.Errorf(got)
					}
				}
			})
		}
	}
}
