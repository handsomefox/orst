package sorts

import (
	"github.com/handsomefox/orst/sorter"
	"golang.org/x/exp/constraints"
)

type (
	Insertion[T constraints.Ordered] struct{}
	InsertionAny[T any]              struct{}
)

func (b Insertion[T]) Sort(s []T, cmp sorter.Comparator[T]) {
	insertionSortImpl(s, cmp)
}

func (b InsertionAny[T]) SortAny(s []T, cmp sorter.Comparator[T]) {
	insertionSortImpl(s, cmp)
}

func insertionSortImpl[T any](s []T, cmp sorter.Comparator[T]) {
	for unsorted := 1; unsorted < len(s); unsorted++ {
		i := unsorted
		for i > 0 && cmp(&s[i], &s[i-1]) {
			s[i-1], s[i] = s[i], s[i-1]
			i--
		}
	}
}
