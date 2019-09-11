package util

import "testing"

var testBytes = []byte("hello world")
var testString = "hello world"

func BenchmarkBtsNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(testBytes)
	}
}

func BenchmarkBtsUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = String(testBytes)
	}
}

func BenchmarkStbNatvie(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte(testString)
	}
}

func BenchmarkStbUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Bytes(testString)
	}
}
