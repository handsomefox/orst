package orst

import (
	"golang.org/x/exp/constraints"
)

type (
	Selection[T constraints.Ordered] struct{}
	SelectionAny[T any]              struct{}
)

func (b Selection[T]) Sort(s []T, cmp Comparator[T]) {
	if cmp == nil { // if no comparison function was specified,
		cmp = func(i, j T) bool { // use the default operator <
			return i < j
		}
	}
	selectionSortImpl(s, cmp)
}

func (b SelectionAny[T]) SortAny(s []T, cmp Comparator[T]) {
	if cmp == nil { // return as there is no way to compare the items
		return
	}
	selectionSortImpl(s, cmp)
}

func selectionSortImpl[T any](s []T, cmp Comparator[T]) {
	for unsorted := range s {
		smallest := unsorted
		for i := unsorted + 1; i < len(s); i++ {
			if cmp(s[i], s[smallest]) {
				smallest = i
			}
		}
		if unsorted != smallest {
			s[unsorted], s[smallest] = s[smallest], s[unsorted]
		}
	}
}
