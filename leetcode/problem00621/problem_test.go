package problem00621

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_maximumGap(t *testing.T) {
	funs := map[string]func([]byte, int) int{
		"leastInterval": leastInterval,
	}

	tests := map[string]struct {
		tasks []byte
		n     int
		want  int
	}{
		"normal1": {
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     2,
			want:  8,
		},
		"normal2": {
			tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:     0,
			want:  6,
		},
		"normal3": {
			tasks: []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'},
			n:     2,
			want:  16,
		},
		"normal4": {
			tasks: []byte{'A', 'B', 'C', 'A', 'B', 'C', 'A', 'B', 'C', 'A', 'B', 'C', 'A', 'B', 'C', 'A', 'B', 'C'},
			n:     2,
			want:  18,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.tasks, tt.n)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
