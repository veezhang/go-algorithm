package problem00127

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_ladderLength(t *testing.T) {
	funs := map[string]func(string, string, []string) int{
		"ladderLength1": ladderLengthBFS1,
		"ladderLength2": ladderLengthBFS2,
		"ladderLength3": ladderLengthBFS3,
	}

	tests := map[string]struct {
		beginWord string
		endWord   string
		wordList  []string
		want      int
	}{
		"normal1": {
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			want:      5,
		},
		"normal2": {
			beginWord: "hit",
			endWord:   "dog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			want:      4,
		},
		"normal3": {
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			want:      0,
		},
		"normal4": {
			beginWord: "a",
			endWord:   "c",
			wordList:  []string{"a", "b", "c"},
			want:      2,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.beginWord, tt.endWord, tt.wordList)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
