package util

import "testing"

var testBytes = []byte("hello world")

func BenchmarkNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(testBytes)
	}
}

func BenchmarkUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = String(testBytes)
	}
}
