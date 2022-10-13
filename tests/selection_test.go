package orst_test

import (
	"testing"

	"github.com/handsomefox/orst"
)

func Benchmark_SelectionSort10(b *testing.B) {
	const sliceSize = 10
	sorter := orst.Selection[int]{}

	b.StopTimer()
	cmpCounter := 0
	cmp := func(i, j int) bool {
		cmpCounter++
		return i < j
	}
	s := generateSlice(sliceSize, true)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.Sort[int](s, sorter, cmp)
	}
	if log {
		b.Log("total comparisons: ", cmpCounter)
	}
}

func Benchmark_SelectionSort100(b *testing.B) {
	const sliceSize = 100
	sorter := orst.Selection[int]{}

	b.StopTimer()
	cmpCounter := 0
	cmp := func(i, j int) bool {
		cmpCounter++
		return i < j
	}
	s := generateSlice(sliceSize, true)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.Sort[int](s, sorter, cmp)
	}
	if log {
		b.Log("total comparisons: ", cmpCounter)
	}
}

func Benchmark_SelectionSort1000(b *testing.B) {
	const sliceSize = 1_000
	sorter := orst.Selection[int]{}

	b.StopTimer()
	cmpCounter := 0
	cmp := func(i, j int) bool {
		cmpCounter++
		return i < j
	}
	s := generateSlice(sliceSize, true)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.Sort[int](s, sorter, cmp)
	}
	if log {
		b.Log("total comparisons: ", cmpCounter)
	}
}

func Benchmark_SelectionSortAny10(b *testing.B) {
	const sliceSize = 10
	sorter := orst.SelectionAny[custom]{}

	b.StopTimer()
	cmpCounter := 0
	cmp := func(i, j custom) bool {
		cmpCounter++
		return i.data < j.data
	}
	s := generateSliceCustom(sliceSize, true)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.SortAny[custom](s, sorter, cmp)
	}
	if log {
		b.Log("total comparisons: ", cmpCounter)
	}
}

func Benchmark_SelectionSortAny100(b *testing.B) {
	const sliceSize = 100
	sorter := orst.SelectionAny[custom]{}

	b.StopTimer()
	cmpCounter := 0
	cmp := func(i, j custom) bool {
		cmpCounter++
		return i.data < j.data
	}
	s := generateSliceCustom(sliceSize, true)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.SortAny[custom](s, sorter, cmp)
	}
	if log {
		b.Log("total comparisons: ", cmpCounter)
	}
}

func Benchmark_SelectionSortAny1000(b *testing.B) {
	const sliceSize = 1_000
	sorter := orst.SelectionAny[custom]{}

	b.StopTimer()
	cmpCounter := 0
	cmp := func(i, j custom) bool {
		cmpCounter++
		return i.data < j.data
	}
	s := generateSliceCustom(sliceSize, true)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		orst.SortAny[custom](s, sorter, cmp)
	}
	if log {
		b.Log("total comparisons: ", cmpCounter)
	}
}
