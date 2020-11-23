package problem00242

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_insertionSortList(t *testing.T) {
	funs := map[string]func(string, string) bool{
		"isAnagram": isAnagram,
	}

	tests := map[string]struct {
		s    string
		t    string
		want bool
	}{
		"normal1": {
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		"normal2": {
			s:    "rat",
			t:    "car",
			want: false,
		},
		"normal3": {
			s:    "Hello 中国！",
			t:    "国中！ Hello",
			want: true,
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
