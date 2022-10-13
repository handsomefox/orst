package sorts_test

import (
	"testing"

	"github.com/handsomefox/orst"
)

func Benchmark_QuickSort10(b *testing.B) {
	b.StopTimer()
	const sliceSize = 10
	s := generateSlice(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.Sort(s, orst.QuickSort, intComparator)
	}
}

func Benchmark_QuickSort100(b *testing.B) {
	b.StopTimer()
	const sliceSize = 100
	s := generateSlice(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.Sort(s, orst.QuickSort, intComparator)
	}
}

func Benchmark_QuickSort1000(b *testing.B) {
	b.StopTimer()
	const sliceSize = 1000
	s := generateSlice(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.Sort(s, orst.QuickSort, intComparator)
	}
}

func Benchmark_QuickSortAny10(b *testing.B) {
	b.StopTimer()
	const sliceSize = 10
	s := generateSliceAny(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.SortAny(s, orst.QuickSort, exampleComparator)
	}
}

func Benchmark_QuickSortAny100(b *testing.B) {
	b.StopTimer()
	const sliceSize = 100
	s := generateSliceAny(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.SortAny(s, orst.QuickSort, exampleComparator)
	}
}

func Benchmark_QuickSortAny1000(b *testing.B) {
	b.StopTimer()
	const sliceSize = 1000
	s := generateSliceAny(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.SortAny(s, orst.QuickSort, exampleComparator)
	}
}
