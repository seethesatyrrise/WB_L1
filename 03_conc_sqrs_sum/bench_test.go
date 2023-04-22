/*
BenchmarkSample1-4         92966             12059 ns/op
BenchmarkSample2-4         57657             20053 ns/op
BenchmarkSample3-4         39262             30682 ns/op
*/

package main

import (
	"testing"
)

func BenchmarkSample1(b *testing.B) {
	input := []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i := 0; i < b.N; i++ {
		calculateSum1(input)
	}
}

func BenchmarkSample2(b *testing.B) {
	input := []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i := 0; i < b.N; i++ {
		calculateSum2(input)
	}
}

func BenchmarkSample3(b *testing.B) {
	input := []int64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i := 0; i < b.N; i++ {
		calculateSum3(input)
	}
}
