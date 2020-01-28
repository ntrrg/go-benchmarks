package recursion_test

import (
	"testing"

	"recursion"
)

func TestFib(t *testing.T) {
	cases := []struct{
		in, want int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
	}

	for _, c := range cases {
		r := recursion.Fib(c.in)
		if r != c.want {
			t.Errorf("Fib(%v) == %v; want %v", c.in, r, c.want)
		}

		r = recursion.FibRec(c.in)
		if r != c.want {
			t.Errorf("FibRec(%v) == %v; want %v", c.in, r, c.want)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		recursion.Fib(40)
	}
}

func BenchmarkFibRec(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		recursion.FibRec(40)
	}
}
