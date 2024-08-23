package main

import "testing"

func BenchmarkConvertWith0Len(b *testing.B) {
	input := []Foo{"1", "2", "3", "4", "5"}
	for i := 0; i < b.N; i++ {
		convertWith0Len(input)
	}
}

func BenchmarkConvertWithLen(b *testing.B) {
	input := []Foo{"1", "2", "3", "4", "5"}
	for i := 0; i < b.N; i++ {
		convertWithLen(input)
	}
}

func BenchmarkConvertWithCap(b *testing.B) {
	input := []Foo{"1", "2", "3", "4", "5"}
	for i := 0; i < b.N; i++ {
		convertWithCap(input)
	}
}
