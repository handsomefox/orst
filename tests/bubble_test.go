package orst_test

import (
	"testing"

	"github.com/handsomefox/orst"
)

func Benchmark_BubbleSort10(b *testing.B) {
	const sliceSize = 10
	sorter := orst.Bubble[int]{}

	b.StopTimer()
	cmpCounter := 0
	cmp := func(i, j int) bool { cmpCounter++; return i < j }
	s := generateSlice(sliceSize, true)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.Sort[int](s, sorter, cmp)
	}
	if log {
		b.Log("total comparisons: ", cmpCounter)
	}
}

func Benchmark_BubbleSort100(b *testing.B) {
	const sliceSize = 100
	sorter := orst.Bubble[int]{}

	b.StopTimer()
	cmpCounter := 0
	cmp := func(i, j int) bool { cmpCounter++; return i < j }
	s := generateSlice(sliceSize, true)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.Sort[int](s, sorter, cmp)
	}
	if log {
		b.Log("total comparisons: ", cmpCounter)
	}
}

func Benchmark_BubbleSort1000(b *testing.B) {
	const sliceSize = 1_000
	sorter := orst.Bubble[int]{}

	b.StopTimer()
	cmpCounter := 0
	cmp := func(i, j int) bool { cmpCounter++; return i < j }
	s := generateSlice(sliceSize, true)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.Sort[int](s, sorter, cmp)
	}
	if log {
		b.Log("total comparisons: ", cmpCounter)
	}
}

func Benchmark_BubbleSortAny10(b *testing.B) {
	const sliceSize = 10
	sorter := orst.BubbleAny[custom]{}

	b.StopTimer()
	cmpCounter := 0
	cmp := func(i, j custom) bool { cmpCounter++; return i.data < j.data }
	s := generateSliceCustom(sliceSize, true)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.SortAny[custom](s, sorter, cmp)
	}
	if log {
		b.Log("total comparisons: ", cmpCounter)
	}
}

func Benchmark_BubbleSortAny100(b *testing.B) {
	const sliceSize = 100
	sorter := orst.BubbleAny[custom]{}

	b.StopTimer()
	cmpCounter := 0
	cmp := func(i, j custom) bool { cmpCounter++; return i.data < j.data }
	s := generateSliceCustom(sliceSize, true)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.SortAny[custom](s, sorter, cmp)
	}
	if log {
		b.Log("total comparisons: ", cmpCounter)
	}
}

func Benchmark_BubbleSortAny1000(b *testing.B) {
	const sliceSize = 1_000
	sorter := orst.BubbleAny[custom]{}

	b.StopTimer()
	cmpCounter := 0
	cmp := func(i, j custom) bool { cmpCounter++; return i.data < j.data }
	s := generateSliceCustom(sliceSize, true)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.SortAny[custom](s, sorter, cmp)
	}
	if log {
		b.Log("total comparisons: ", cmpCounter)
	}
}
