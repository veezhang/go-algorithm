package problem80108

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_SetZeroes(t *testing.T) {
	funs := map[string]func(matrix [][]int){
		"setZeroes": setZeroes,
	}

	tests := map[string]struct {
		matrix [][]int
		want   [][]int
	}{
		"normal1": {
			matrix: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			want: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		"normal2": {
			matrix: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
		"empty1": {
			matrix: [][]int{
				{},
			},
			want: [][]int{
				{},
			},
		},
		"empty2": {
			matrix: [][]int{
				{},
				{},
			},
			want: [][]int{
				{},
				{},
			},
		},
		"empty3": {
			matrix: [][]int{},
			want:   [][]int{},
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				fun(tt.matrix)
				got := tt.matrix
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
