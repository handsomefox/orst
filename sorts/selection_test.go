package sorts_test

import (
	"testing"

	"github.com/handsomefox/orst"
)

func Benchmark_SelectionSort10(b *testing.B) {
	b.StopTimer()
	const sliceSize = 10
	s := generateSlice(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.Sort(s, orst.SelectionSort, intComparator)
	}
}

func Benchmark_SelectionSort100(b *testing.B) {
	b.StopTimer()
	const sliceSize = 100
	s := generateSlice(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.Sort(s, orst.SelectionSort, intComparator)
	}
}

func Benchmark_SelectionSort1000(b *testing.B) {
	b.StopTimer()
	const sliceSize = 1000
	s := generateSlice(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.Sort(s, orst.SelectionSort, intComparator)
	}
}

func Benchmark_SelectionSortAny10(b *testing.B) {
	b.StopTimer()
	const sliceSize = 10
	s := generateSliceAny(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.SortAny(s, orst.SelectionSort, exampleComparator)
	}
}

func Benchmark_SelectionSortAny100(b *testing.B) {
	b.StopTimer()
	const sliceSize = 100
	s := generateSliceAny(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.SortAny(s, orst.SelectionSort, exampleComparator)
	}
}

func Benchmark_SelectionSortAny1000(b *testing.B) {
	b.StopTimer()
	const sliceSize = 1000
	s := generateSliceAny(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.SortAny(s, orst.SelectionSort, exampleComparator)
	}
}
