package problem00006

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Convert(t *testing.T) {
	tests := map[string]struct {
		s       string
		numRows int
		want    string
	}{
		"normal": {
			s:       "LEETCODEISHIRING",
			numRows: 3,
			want:    "LCIRETOESIIGEDHN",
		},
		"empty": {
			s:       "",
			numRows: 3,
			want:    "",
		},
		"less then row": {
			s:       "AB",
			numRows: 3,
			want:    "AB",
		},
		"eq the row": {
			s:       "ABC",
			numRows: 3,
			want:    "ABC",
		},
		"one row": {
			s:       "ABC",
			numRows: 1,
			want:    "ABC",
		},
	}

	// 测试ListNode.ToSlice
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := convert(tt.s, tt.numRows)
			want := tt.want
			diff := cmp.Diff(got, want)
			if diff != "" {
				t.Errorf(diff)
			}
		})
	}
}
