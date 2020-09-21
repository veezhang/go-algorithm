package problem00240

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_SearchMatrix(t *testing.T) {
	funs := map[string]func(matrix [][]int, target int) bool{
		"searchMatrixStep":      searchMatrixStep,
		"searchMatrixSubMatrix": searchMatrixSubMatrix,
	}

	tests := map[string]struct {
		matrix [][]int
		target int
		want   bool
	}{
		"normal1": {
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 5,
			want:   true,
		},
		"normal2": {
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 20,
			want:   false,
		},
		"normal3": {
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 13,
			want:   true,
		},
		"normal4": {
			matrix: [][]int{
				{1, 4, 7, 11, 15},
				{2, 5, 8, 12, 19},
				{3, 6, 9, 16, 22},
				{10, 13, 14, 17, 24},
				{18, 21, 23, 26, 30},
			},
			target: 18,
			want:   true,
		},
		"empty1": {
			matrix: [][]int{
				{},
			},
			target: 0,
			want:   false,
		},
		"empty2": {
			matrix: [][]int{
				{},
				{},
			},
			target: 0,
			want:   false,
		},
		"empty3": {
			matrix: [][]int{},
			target: 0,
			want:   false,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.matrix, tt.target)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
