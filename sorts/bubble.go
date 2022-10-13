package sorts

import (
	"github.com/handsomefox/orst/sorter"
	"golang.org/x/exp/constraints"
)

type (
	Bubble[T constraints.Ordered] struct{}
	BubbleAny[T any]              struct{}
)

func (b Bubble[T]) Sort(s []T, cmp sorter.Comparator[T]) {
	bubbleSortImpl(s, cmp)
}

func (b BubbleAny[T]) SortAny(s []T, cmp sorter.Comparator[T]) {
	bubbleSortImpl(s, cmp)
}

func bubbleSortImpl[T any](s []T, cmp sorter.Comparator[T]) {
	swapped := true
	for swapped {
		swapped = false
		for i := range s {
			if i == len(s)-1 {
				continue
			}
			if !cmp(s[i], s[i+1]) {
				s[i], s[i+1] = s[i+1], s[i]
				swapped = true
			}
		}
	}
}
