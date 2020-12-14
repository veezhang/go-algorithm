package problem00842

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_groupAnagrams(t *testing.T) {
	funs := map[string]func([]string) [][]string{
		"groupAnagrams1": groupAnagrams1,
		"groupAnagrams2": groupAnagrams2,
	}

	tests := map[string]struct {
		strs []string
		want [][]string
	}{
		"normal1": {
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"ate", "eat", "tea"},
				{"nat", "tan"},
				{"bat"},
			},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.strs)
				want := tt.want

				sortfn := func(strss [][]string) {
					for _, strs := range strss {
						sort.Strings(strs)
					}
					sort.Slice(strss, func(i, j int) bool {
						return strss[i][0] < strss[j][0]
					})
				}

				sortfn(got)
				sortfn(want)

				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
