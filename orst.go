// Package orst is a package that has different sorting methods
// that can be applied to generic slices or with your own comparator function
package orst

import "golang.org/x/exp/constraints"

// Sort sorts the given slice (by mutating it) with the specified Sorter.
// If less is nil, sorts in ascending order using the inbuilt comparison operators.
func Sort[T constraints.Ordered](s []T, srt Sorter[T], cmp Comparator[T]) {
	srt.Sort(s, cmp)
}

// Sort sorts the given slice (by mutating it) with the specified Sorter.
// If less is nil, the function immediately returns.
func SortAny[T any](s []T, srt SorterAny[T], cmp Comparator[T]) {
	srt.SortAny(s, cmp)
}
