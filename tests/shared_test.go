package orst_test

type custom struct {
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

func generateSliceCustom(count int, reversed bool) []custom {
	s := make([]custom, 0, count)
	if !reversed {
		for i := 0; i < count; i++ {
			s = append(s, custom{i})
		}
	} else {
		for i := count - 1; i > -1; i-- {
			s = append(s, custom{i})
		}
	}
	return s
}
