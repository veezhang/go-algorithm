package problem00714

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_findTheDifference(t *testing.T) {
	funs := map[string]func(s string, t string) byte{
		"findTheDifference1": findTheDifference1,
		"findTheDifference2": findTheDifference2,
	}

	tests := map[string]struct {
		s    string
		t    string
		want byte
	}{
		"normal1": {
			s:    "abcd",
			t:    "abcde",
			want: 'e',
		},
		"normal2": {
			s:    "",
			t:    "y",
			want: 'y',
		},
		"normal3": {
			s:    "a",
			t:    "aa",
			want: 'a',
		},
		"normal4": {
			s:    "ae",
			t:    "aea",
			want: 'a',
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.s, tt.t)
				want := tt.want

				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
