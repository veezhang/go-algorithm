package problem00028

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_StrStr(t *testing.T) {
	funs := map[string]func(haystack string, needle string) int{
		"strStrSubstring":            strStrSubstring,
		"strStrTwoPointer":           strStrTwoPointer,
		"strStrKMP":                  strStrKMP,
		"strStrBM":                   strStrBM,
		"strStrHorspool":             strStrHorspool,
		"strStrSunday":               strStrSunday,
		"strStrSundayWithShiftTable": strStrSundayWithShiftTable,
		"strStrRabinKarp":            strStrRabinKarp,
		"strStrGO":                   strStrGO,
	}

	tests := map[string]struct {
		haystack string
		needle   string
		want     int
	}{
		"normal": {
			haystack: "hello",
			needle:   "ll",
			want:     2,
		},
		"not found": {
			haystack: "aaaaa",
			needle:   "bba",
			want:     -1,
		},
		"endmartch": {
			haystack: "hello",
			needle:   "lo",
			want:     3,
		},
		"mississippi": {
			haystack: "mississippi",
			needle:   "issip",
			want:     4,
		},
		"AABADAAABBCCCADABCBCDA": {
			haystack: "AABADAAABBCCCADABCBCDA",
			needle:   "ABCBC",
			want:     15,
		},
		"BCABCDABCEABFBCABCDABCEAABCBCBCEABCBCABCDCBCABCDABCEABCAA": {
			haystack: "BCABCDABCEABFBCABCDABCEAABCBCBCEABCBCABCDCBCABCDABCEABCAA",
			needle:   "BCABCDABCEABC",
			want:     42,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.haystack, tt.needle)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
