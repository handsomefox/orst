package sorts_test

import (
	"testing"

	"github.com/handsomefox/orst"
)

func Benchmark_InsertionSort10(b *testing.B) {
	b.StopTimer()
	const sliceSize = 10
	s := generateSlice(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.Sort(s, orst.InsertionSort, intComparator)
	}
}

func Benchmark_InsertionSort100(b *testing.B) {
	b.StopTimer()
	const sliceSize = 100
	s := generateSlice(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.Sort(s, orst.InsertionSort, intComparator)
	}
}

func Benchmark_InsertionSort1000(b *testing.B) {
	b.StopTimer()
	const sliceSize = 1000
	s := generateSlice(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.Sort(s, orst.InsertionSort, intComparator)
	}
}

func Benchmark_InsertionSortAny10(b *testing.B) {
	b.StopTimer()
	const sliceSize = 10
	s := generateSliceAny(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.SortAny(s, orst.InsertionSort, exampleComparator)
	}
}

func Benchmark_InsertionSortAny100(b *testing.B) {
	b.StopTimer()
	const sliceSize = 100
	s := generateSliceAny(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.SortAny(s, orst.InsertionSort, exampleComparator)
	}
}

func Benchmark_InsertionSortAny1000(b *testing.B) {
	b.StopTimer()
	const sliceSize = 1000
	s := generateSliceAny(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.SortAny(s, orst.InsertionSort, exampleComparator)
	}
}
