package problem00003

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	funs := map[string]func(s string) int{
		"lengthOfLongestSubstring": lengthOfLongestSubstring,
	}

	tests := map[string]struct {
		s    string
		want int
	}{
		"abcabcbb": {
			s:    "abcabcbb",
			want: 3,
		},
		"bbbbb": {
			s:    "bbbbb",
			want: 1,
		},
		"pwwkew": {
			s:    "pwwkew",
			want: 3,
		},
		"empty": {
			s:    "",
			want: 0,
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
