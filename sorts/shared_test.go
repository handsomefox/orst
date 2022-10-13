package sorts_test

const log = false

var (
	intComparator     = func(i, j int) bool { return i < j }
	exampleComparator = func(i, j exampleType) bool { return i.data < j.data }
)

type exampleType struct {
	data int
}

func generateSlice(count int, reversed bool) []int {
	s := make([]int, 0, count)
	if !reversed {
		for i := 0; i < count; i++ {
			s = append(s, i)
		}
	} else {
		for i := count - 1; i > -1; i-- {
			s = append(s, i)
		}
	}
	return s
}

func generateSliceAny(count int, reversed bool) []exampleType {
	s := make([]exampleType, 0, count)
	if !reversed {
		for i := 0; i < count; i++ {
			s = append(s, exampleType{i})
		}
	} else {
		for i := count - 1; i > -1; i-- {
			s = append(s, exampleType{i})
		}
	}
	return s
}
