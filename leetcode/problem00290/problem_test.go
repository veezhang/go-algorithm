package problem00290

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_wordPattern(t *testing.T) {
	funs := map[string]func(string, s string) bool{
		"wordPattern": wordPattern,
	}

	tests := map[string]struct {
		pattern string
		s       string
		want    bool
	}{
		"normal1": {
			pattern: "abba",
			s:       "dog cat cat dog",
			want:    true,
		},
		"normal2": {
			pattern: "abba",
			s:       "dog cat cat fish",
			want:    false,
		},
		"normal3": {
			pattern: "aaaa",
			s:       "dog cat cat dog",
			want:    false,
		},
		"normal4": {
			pattern: "abba",
			s:       "dog dog dog dog",
			want:    false,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.pattern, tt.s)
				want := tt.want

				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
