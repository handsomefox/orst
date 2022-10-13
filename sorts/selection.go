package sorts

import (
	"github.com/handsomefox/orst/sorter"
	"golang.org/x/exp/constraints"
)

type (
	Selection[T constraints.Ordered] struct{}
	SelectionAny[T any]              struct{}
)

func (b Selection[T]) Sort(s []T, cmp sorter.Comparator[T]) {
	selectionSortImpl(s, cmp)
}

func (b SelectionAny[T]) SortAny(s []T, cmp sorter.Comparator[T]) {
	selectionSortImpl(s, cmp)
}

func selectionSortImpl[T any](s []T, cmp sorter.Comparator[T]) {
	for unsorted := range s {
		smallest := unsorted
		for i := unsorted + 1; i < len(s); i++ {
			if cmp(&s[i], &s[smallest]) {
				smallest = i
			}
		}
		if unsorted != smallest {
			s[unsorted], s[smallest] = s[smallest], s[unsorted]
		}
	}
}
