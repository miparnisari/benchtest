package fib

import (
	"testing"
)

func BenchmarkFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Fib(13)
	}
}

func BenchmarkFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Fib(23)
	}
}

func BenchmarkFib25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Fib(30)
	}
}
