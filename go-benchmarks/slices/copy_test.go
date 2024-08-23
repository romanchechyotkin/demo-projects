package slices

import "testing"

func BenchmarkCopySlice(b *testing.B) {
	input := []string{"1", "2", "3", "4", "5"}
	for i := 0; i < b.N; i++ {
		CopySlice(input)
	}
}
