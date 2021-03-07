package pkg_test

import (
	"testing"

	"pkg"
)

const N = 100

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]pkg.Tree, N)

		for j := 0; j < N; j++ {
			t := &s[j]
			t.SetName("Miguel Angel")
			t.HasChildren()
		}
	}
}

func BenchmarkSliceOfPointers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]*pkg.Tree, 0, N)

		for j := 0; j < N; j++ {
			s = append(s, &pkg.Tree{})
		}

		for j := 0; j < N; j++ {
			t := s[j]
			t.SetName("Miguel Angel")
			t.HasChildren()
		}
	}
}

func BenchmarkSliceToSliceOfPointers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]pkg.Tree, N)
		sp := make([]*pkg.Tree, 0, N)

		for j := 0; j < N; j++ {
			sp = append(sp, &s[j])
		}

		for j := 0; j < N; j++ {
			t := sp[j]
			t.SetName("Miguel Angel")
			t.HasChildren()
		}
	}
}

func BenchmarkChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := make(chan pkg.Tree, N)

		for j := 0; j < N; j++ {
			c <- pkg.Tree{}
		}

		for j := 0; j < N; j++ {
			x := <-c
			t := &x
			t.SetName("Miguel Angel")
			t.HasChildren()
		}
	}
}

func BenchmarkChannelOfPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := make(chan *pkg.Tree, N)

		for j := 0; j < N; j++ {
			c <- &pkg.Tree{}
		}

		for j := 0; j < N; j++ {
			t := <-c
			t.SetName("Miguel Angel")
			t.HasChildren()
		}
	}
}
