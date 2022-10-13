# orst

This is a package that has different sorting algorithms that can be applied to generic slices or used with your own comparator function.

## Usage example

### With a normal slice

```go
slice := []int{3,2,1}  // slice = (3,2,1)
// with default comparator
orst.Sort(slice, orst.BubbleSort, nil) // slice = (1,2,3)
// or with custom comparator
orst.Sort(slice, orst.BubbleSort, func(i,j *int) bool { return *j < *i }) // slice = (3,2,1)
```

### With a custom type

```go
type example struct { val int }
slice := []example{{3},{2},{1}} // slice = (3,2,1)
// only with custom comparator
orst.SortAny(slice, orst.BubbleSort, func(i,j *example) bool { return i.val < j.val }) // slice = (1,2,3)
```

## Currently implemented sorting algorithms

1. Bubblesort
2. Insertion sort
3. Quicksort
4. Selection sort
