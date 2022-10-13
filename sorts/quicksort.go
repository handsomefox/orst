package sorts

import (
	"math/rand"
	"time"

	"github.com/handsomefox/orst/sorter"
	"golang.org/x/exp/constraints"
)

type (
	Quicksort[T constraints.Ordered] struct{}
	QuicksortAny[T any]              struct{}
)

func (b Quicksort[T]) Sort(s []T, cmp sorter.Comparator[T]) {
	quicksortSortImpl(s, cmp)
}

func (b QuicksortAny[T]) SortAny(s []T, cmp sorter.Comparator[T]) {
	quicksortSortImpl(s, cmp)
}

func quicksortSortImpl[T any](s []T, cmp sorter.Comparator[T]) {
	if len(s) == 1 || len(s) == 0 {
		return
	} else if len(s) == 2 {
		if cmp(s[1], s[0]) {
			s[0], s[1] = s[1], s[0]
			return
		}
	}

	left, right := 0, len(s)-1

	rand.Seed(time.Now().UnixNano())
	pivot := rand.Int() % len(s)
	s[pivot], s[right] = s[right], s[pivot]

	for i := range s {
		if cmp(s[i], s[right]) {
			s[left], s[i] = s[i], s[left]
			left++
		}
	}

	s[left], s[right] = s[right], s[left]
	quicksortSortImpl(s[:left], cmp)
	quicksortSortImpl(s[left+1:], cmp)
}
