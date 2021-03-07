package main

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", 42)
	}
}

func BenchmarkFmtV(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%v", 42)
	}
}

func BenchmarkStrconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(42)
	}
}
