package problem00005

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_LongestPalindrome(t *testing.T) {
	funs := map[string]func(string) string{
		"longestPalindromeLongestSubstring": longestPalindromeLongestSubstring,
		"longestPalindromeDP":               longestPalindromeDP,
		"longestPalindromeExandCenter":      longestPalindromeExandCenter,
		"longestPalindromeManacher":         longestPalindromeManacher,
	}

	tests := map[string]struct {
		s    string
		want string
	}{
		"normal": {
			s:    "babad",
			want: "bab",
		},
		"long": {
			s:    "abcacdedcaefasdeeeedsavasderasdadsad",
			want: "asdeeeedsa",
		},
		"empty": {
			s:    "",
			want: "",
		},
		"1char": {
			s:    "a",
			want: "a",
		},
		"2samechar": {
			s:    "aa",
			want: "aa",
		},
		"2char": {
			s:    "ab",
			want: "a",
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
