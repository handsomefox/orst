package orst

import "golang.org/x/exp/constraints"

type (
	Bubble[T constraints.Ordered] struct{}
	BubbleAny[T any]              struct{}
)

func (b Bubble[T]) Sort(s []T, cmp Comparator[T]) {
	if cmp == nil { // if no comparison function was specified,
		cmp = func(i, j T) bool { // use the default operator <
			return i < j
		}
	}
	bubbleSortImpl(s, cmp)
}

func (b BubbleAny[T]) SortAny(s []T, cmp Comparator[T]) {
	if cmp == nil { // return as there is no way to compare the items
		return
	}
	bubbleSortImpl(s, cmp)
}

func bubbleSortImpl[T any](s []T, cmp Comparator[T]) {
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
