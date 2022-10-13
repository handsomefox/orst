package orst_test

import (
	"testing"

	"github.com/handsomefox/orst"
)

const log = false

func Test_Sort(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		sorter orst.Sorter[int]
		desc   string
	}{
		{sorter: orst.Bubble[int]{}, desc: "BubbleSort"},
		{sorter: orst.Insertion[int]{}, desc: "InsertionSort"},
	}

	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			s, s2 := generateSlice(250, true), generateSlice(250, false)

			tC.sorter.Sort(s, nil)

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
		sorter orst.SorterAny[custom]
		desc   string
	}{
		{sorter: orst.BubbleAny[custom]{}, desc: "BubbleSortAny"},
		{sorter: orst.InsertionAny[custom]{}, desc: "InsertionSort"},
	}

	var cmp orst.Comparator[custom] = func(i, j custom) bool {
		return i.data < j.data
	}

	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			s, s2 := generateSliceCustom(250, true), generateSliceCustom(250, false)

			tC.sorter.SortAny(s, cmp)

			for i := 0; i < len(s); i++ {
				if s[i].data != s2[i].data {
					t.Fatalf("%s failed, expected: %v, got: %v", tC.desc, s2[i], s[i])
				}
			}
		})
	}
}
