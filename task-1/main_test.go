package main

import "testing"

func Benchmark_doubles(b *testing.B) {
	array := []int{1, 3, 1, 3, 1, 3, 8, 8, 3, 8}
	k := 2

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		doubles(array, k)
	}
}
