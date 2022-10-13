package sorter

import "golang.org/x/exp/constraints"

// Comparator should return if i is less than j.
type Comparator[T any] func(i, j *T) bool

// Sorter is the interface that has to be implemented by sorting algorithms in this package.
type Sorter[T constraints.Ordered] interface {
	Sort([]T, Comparator[T])
}

// SorterAny is the interface that has to be implemented by sorting algorithms in this package
// that want to sort any slice type.
type SorterAny[T any] interface {
	SortAny([]T, Comparator[T])
}
