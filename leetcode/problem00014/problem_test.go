package problem00014

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_LongestCommonPrefix(t *testing.T) {
	funs := map[string]func(strs []string) string{
		"longestCommonPrefix": longestCommonPrefix,
	}

	tests := map[string]struct {
		strs []string
		want string
	}{
		"normal": {
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		"empty strs": {
			strs: []string{},
			want: "",
		},
		"with empty": {
			strs: []string{"flower", "flow", ""},
			want: "",
		},
		"1th not equal": {
			strs: []string{"flower", "flow", "olight"},
			want: "",
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.strs)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
