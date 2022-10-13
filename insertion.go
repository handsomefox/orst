package orst

import "golang.org/x/exp/constraints"

type (
	Insertion[T constraints.Ordered] struct{}
	InsertionAny[T any]              struct{}
)

func (b Insertion[T]) Sort(s []T, cmp Comparator[T]) {
	if cmp == nil { // if no comparison function was specified,
		cmp = func(i, j T) bool { // use the default operator <
			return i < j
		}
	}
	insertionSortImpl(s, cmp)
}

func (b InsertionAny[T]) SortAny(s []T, cmp Comparator[T]) {
	if cmp == nil { // return as there is no way to compare the items
		return
	}
	insertionSortImpl(s, cmp)
}

func insertionSortImpl[T any](s []T, cmp Comparator[T]) {
	for unsorted := 1; unsorted < len(s); unsorted++ {
		i := unsorted
		for i > 0 && cmp(s[i], s[i-1]) {
			s[i-1], s[i] = s[i], s[i-1]
			i--
		}
	}
}
