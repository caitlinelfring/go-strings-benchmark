package main

import (
	"testing"
)

func BenchmarkBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Builder()
	}
}

func BenchmarkBuilderWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuilderWithCap()
	}
}

func BenchmarkPlusEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PlusEqual()
	}
}
func BenchmarkFmtSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FmtSprint()
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Join()
	}
}
