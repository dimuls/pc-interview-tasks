package main

import "testing"

func Benchmark_mostCommonPrefix(b *testing.B) {
	ss := []string{"qwe", "qweasd", "qwsdfsdf"}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mostCommonPrefix(ss)
	}
}
