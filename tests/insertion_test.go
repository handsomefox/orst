package orst_test

import (
	"testing"

	"github.com/handsomefox/orst"
)

func Benchmark_InsertionSort10(b *testing.B) {
	const sliceSize = 10
	sorter := orst.Insertion[int]{}

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

func Benchmark_InsertionSort100(b *testing.B) {
	const sliceSize = 100
	sorter := orst.Insertion[int]{}

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

func Benchmark_InsertionSort1000(b *testing.B) {
	const sliceSize = 1_000
	sorter := orst.Insertion[int]{}

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

func Benchmark_InsertionSortAny10(b *testing.B) {
	const sliceSize = 10
	sorter := orst.InsertionAny[custom]{}

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

func Benchmark_InsertionSortAny100(b *testing.B) {
	const sliceSize = 100
	sorter := orst.InsertionAny[custom]{}

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

func Benchmark_InsertionSortAny1000(b *testing.B) {
	const sliceSize = 1_000
	sorter := orst.InsertionAny[custom]{}

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
