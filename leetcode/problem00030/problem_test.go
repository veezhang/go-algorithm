package problem00030

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_FindSubstring(t *testing.T) {
	funs := map[string]func(s string, words []string) []int{
		"findSubstringLoop":          findSubstringLoop,
		"findSubstringSlidingWindow": findSubstringSlidingWindow,
	}

	tests := map[string]struct {
		s     string
		words []string
		want  []int
	}{
		"normal": {
			s:     "barfoothefoobarman",
			words: []string{"foo", "bar"},
			want:  []int{0, 9},
		},
		"repeatwords": {
			s:     "wordgoodgoodgoodbestword",
			words: []string{"word", "good", "best", "good"},
			want:  []int{8},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.s, tt.words)
				want := tt.want
				sort.Sort(sort.IntSlice(got))
				sort.Sort(sort.IntSlice(want))
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
