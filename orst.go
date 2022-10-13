// Package orst is a package that has different sorting methods
// that can be applied to generic slices or with your own comparator function
package orst

import "golang.org/x/exp/constraints"

// Sort sorts the given slice (by mutating it) with the specified Sorter.
// If Comparator[T] is nil, sorts in ascending order using the inbuilt comparison operator (<).
//
// # Example
//
//	slice := []int{3,2,1} // slice = (3,2,1)
//	var sorter orst.Bubble[int]
//	orst.Sort[int](slice, sorter, nil) // slice = (1,2,3)
//
// or with custom comparator
//
//	orst.Sort[int](slice, sorter, func(i,j int) bool { return j > i })
func Sort[T constraints.Ordered](s []T, srt Sorter[T], cmp Comparator[T]) {
	srt.Sort(s, cmp)
}

// Sort sorts the given slice (by mutating it) with the specified Sorter.
// If Comparator[T] is nil, the function immediately returns without any changes to the slice.
//
// # Example
//
//	type example struct { val int }
//	slice := []example{{3},{2},{1}} // slice = (3,2,1)
//	var sorter orst.BubbleAny[example]
//	orst.SortAny[example](slice, sorter, func(i,j example) bool { return i.val < j.val }) // slice = (1,2,3)
func SortAny[T any](s []T, srt SorterAny[T], cmp Comparator[T]) {
	srt.SortAny(s, cmp)
}
