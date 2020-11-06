package problem70015

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_ladderLength(t *testing.T) {
	funs := map[string]func(uint32) int{
		"hammingWeight1":  hammingWeight1,
		"hammingWeight2":  hammingWeight2,
		"hammingWeight3":  hammingWeight3,
		"hammingWeight4":  hammingWeight4,
		"hammingWeight5":  hammingWeight5,
		"hammingWeight6":  hammingWeight6,
		"hammingWeight7":  hammingWeight7,
		"hammingWeight8":  hammingWeight8,
		"hammingWeight9":  hammingWeight9,
		"hammingWeight10": hammingWeight10,
		"hammingWeight11": hammingWeight11,
	}

	tests := map[string]struct {
		num  uint32
		want int
	}{
		"example": {
			num:  0b11011110111001110111001100010101,
			want: 20,
		},

		// "normal1": {
		// 	num:  0,
		// 	want: 0,
		// },
		// "normal2": {
		// 	num:  1,
		// 	want: 1,
		// },
		// "normal3": {
		// 	num:  math.MaxUint32,
		// 	want: 32,
		// },
		// "normal4": {
		// 	num:  math.MaxUint32 - 1,
		// 	want: 31,
		// },
		// "normal5": {
		// 	num:  11,
		// 	want: 3,
		// },
		// "normal6": {
		// 	num:  128,
		// 	want: 1,
		// },
		// "normal7": {
		// 	num:  4294967293,
		// 	want: 31,
		// },
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.num)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}
