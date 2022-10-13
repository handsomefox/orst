package sorts_test

import (
	"testing"

	"github.com/handsomefox/orst"
)

func Test_Sort(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		algorithm orst.Kind
		desc      string
	}{
		{algorithm: orst.BubbleSort, desc: "BubbleSort"},
		{algorithm: orst.InsertionSort, desc: "InsertionSort"},
		{algorithm: orst.SelectionSort, desc: "SelectionSort"},
		{algorithm: orst.QuickSort, desc: "Quicksort"},
	}

	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			s, s2 := generateSlice(1000, true), generateSlice(1000, false)

			orst.Sort(s, tC.algorithm, nil)

			for i := 0; i < len(s); i++ {
				if s[i] != s2[i] {
					t.Fatalf("%s failed, expected: %v, got: %v", tC.desc, s2[i], s[i])
				}
			}
		})
	}
}

func Test_SortAny(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		algorithm orst.Kind
		desc      string
	}{
		{algorithm: orst.BubbleSort, desc: "BubbleSortAny"},
		{algorithm: orst.InsertionSort, desc: "InsertionSortAny"},
		{algorithm: orst.SelectionSort, desc: "SelectionSortAny"},
		{algorithm: orst.QuickSort, desc: "QuicksortAny"},
	}

	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			s, s2 := generateSliceAny(1000, true), generateSliceAny(1000, false)

			orst.SortAny(s, tC.algorithm, exampleComparator)

			for i := 0; i < len(s); i++ {
				if s[i].data != s2[i].data {
					t.Fatalf("%s failed, expected: %v, got: %v", tC.desc, s2[i], s[i])
				}
			}
		})
	}
}
