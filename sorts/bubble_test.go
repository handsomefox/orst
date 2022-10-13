package sorts_test

import (
	"testing"

	"github.com/handsomefox/orst"
)

func Benchmark_BubbleSort10(b *testing.B) {
	b.StopTimer()
	const sliceSize = 10
	s := generateSlice(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.Sort(s, orst.BubbleSort, intComparator)
	}
}

func Benchmark_BubbleSort100(b *testing.B) {
	b.StopTimer()
	const sliceSize = 100
	s := generateSlice(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.Sort(s, orst.BubbleSort, intComparator)
	}
}

func Benchmark_BubbleSort1000(b *testing.B) {
	b.StopTimer()
	const sliceSize = 1000
	s := generateSlice(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.Sort(s, orst.BubbleSort, intComparator)
	}
}

func Benchmark_BubbleSortAny10(b *testing.B) {
	b.StopTimer()
	const sliceSize = 10
	s := generateSliceAny(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.SortAny(s, orst.BubbleSort, exampleComparator)
	}
}

func Benchmark_BubbleSortAny100(b *testing.B) {
	b.StopTimer()
	const sliceSize = 100
	s := generateSliceAny(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.SortAny(s, orst.BubbleSort, exampleComparator)
	}
}

func Benchmark_BubbleSortAny1000(b *testing.B) {
	b.StopTimer()
	const sliceSize = 1000
	s := generateSliceAny(sliceSize, true)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.SortAny(s, orst.BubbleSort, exampleComparator)
	}
}
