// Package orst is a package that has different sorting methods
// that can be applied to generic slices or with your own comparator function
package orst

import (
	"github.com/handsomefox/orst/sorter"
	"github.com/handsomefox/orst/sorts"
	"golang.org/x/exp/constraints"
)

// Describes which sorting algorithm to use.
type Kind uint8

const (
	_ Kind = iota
	BubbleSort
	InsertionSort
	SelectionSort
	QuickSort
)

// Sort sorts the given slice (by mutating it) with the specified Sorter.
// If Comparator[T] is nil, sorts in ascending order using the inbuilt comparison operator (<).
//
// # Example
//
//	slice := []int{3,2,1} // slice = (3,2,1)
//
// with defalt comparator
//
//	orst.Sort(slice, orst.BubbleSort, nil) // slice = (1,2,3)
//
// or with custom comparator
//
//	orst.Sort(slice, orst.BubbleSort, func(i,j int) bool { return j > i })
func Sort[T constraints.Ordered](s []T, algorithm Kind, cmp sorter.Comparator[T]) {
	if s == nil {
		return
	}
	if len(s) == 0 || len(s) == 1 {
		return
	}

	if cmp == nil { // if no comparison function was specified,
		cmp = func(i, j T) bool { // use the default operator <
			return i < j
		}
	}

	var srt sorter.Sorter[T] = nil

	switch algorithm {
	case BubbleSort:
		srt = sorts.Bubble[T]{}
	case InsertionSort:
		srt = sorts.Insertion[T]{}
	case SelectionSort:
		srt = sorts.Selection[T]{}
	case QuickSort:
		srt = sorts.Quicksort[T]{}
	default:
		srt = sorts.Quicksort[T]{}
	}

	srt.Sort(s, cmp)
}

// Sort sorts the given slice (by mutating it) with the specified Sorter.
// If Comparator[T] is nil, the function immediately returns without any changes to the slice.
//
// # Example
//
//	type example struct { val int }
//	slice := []example{{3},{2},{1}} // slice = (3,2,1)
//
// only with custom comparator
//
//	orst.SortAny(slice, orst.BubbleSort, func(i,j example) bool { return i.val < j.val }) // slice = (1,2,3)
func SortAny[T any](s []T, algorithm Kind, cmp sorter.Comparator[T]) {
	if s == nil {
		return
	}
	if len(s) == 0 || len(s) == 1 {
		return
	}

	if cmp == nil { // if no comparison function was specified,
		return // don't do anything
	}

	var srt sorter.SorterAny[T] = nil

	switch algorithm {
	case BubbleSort:
		srt = sorts.BubbleAny[T]{}
	case InsertionSort:
		srt = sorts.InsertionAny[T]{}
	case SelectionSort:
		srt = sorts.SelectionAny[T]{}
	case QuickSort:
		srt = sorts.QuicksortAny[T]{}
	default:
		srt = sorts.QuicksortAny[T]{}
	}

	srt.SortAny(s, cmp)
}
