package bitwise_test

import (
	"testing"

	"bitwise"
)

func TestPow(t *testing.T) {
	cases := []struct{
		in, want int
	}{
		{1, 2},
		{2, 4},
		{3, 8},
		{4, 16},
		{5, 32},
		{6, 64},
		{7, 128},
		{8, 256},
		{9, 512},
	}

	for _, c := range cases {
		r := bitwise.BinPow(c.in)
		if r != c.want {
			t.Errorf("BinPow(%v) == %v; want %v", c.in, r, c.want)
		}

		r = bitwise.Pow(c.in)
		if r != c.want {
			t.Errorf("Pow(%v) == %v; want %v", c.in, r, c.want)
		}
	}
}

func BenchmarkBinPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitwise.BinPow(62)
	}
}

func BenchmarkPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitwise.Pow(62)
	}
}
