package problem00020

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_IsValid(t *testing.T) {
	funs := map[string]func(s string) bool{
		"isValid": isValid,
	}

	tests := map[string]struct {
		s    string
		want bool
	}{
		"normal": {
			s:    "()[]{}",
			want: true,
		},
		"empty": {
			s:    "",
			want: true,
		},
		"{]": {
			s:    "{]",
			want: false,
		},
	}

	for fname, fun := range funs {
		for name, tt := range tests {
			runnme := fmt.Sprintf("[%s]%s", fname, name)
			t.Run(runnme, func(t *testing.T) {
				got := fun(tt.s)
				want := tt.want
				diff := cmp.Diff(got, want)
				if diff != "" {
					t.Errorf(diff)
				}
			})
		}
	}
}

// go test --run=^Benchmark_ByteMapSliceSwitch -bench=. ./leetcode/problem00020/...
// Benchmark_ByteMapSliceSwitch/map-4              59312358                19.3  ns/op
// Benchmark_ByteMapSliceSwitch/slice-4            746922408                1.53 ns/op
// Benchmark_ByteMapSliceSwitch/switch-4           465211090                2.51 ns/op

func Benchmark_ByteMapSliceSwitch(b *testing.B) {
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	keys := []byte{'(', '[', '{'}

	b.Run("map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _, ok := m[keys[i%3]]; ok {
			}
		}
	})

	s := []byte{
		')' - '(': '(',
		']' - '(': '[',
		'}' - '(': '{',
	}

	b.Run("slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _ = s[keys[i%3]-'(']; true {
			}
		}
	})

	sw := func(b byte) byte {
		switch b {
		case '(':
			return ')'
		case '[':
			return ']'
		case '{':
			return '}'
		}
		return '0'
	}

	b.Run("switch", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if _ = sw(keys[i%3]); true {
			}
		}
	})
}
