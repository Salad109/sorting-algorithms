package algorithms

// Sorter defines the behavior of a sorting algorithm
type Sorter interface {
	Sort(arr []int32)
	Name() string
}

// FloatSorter defines the behavior of a sorting algorithm for float32 arrays
type FloatSorter interface {
	SortFloat(arr []float32)
	Name() string
}
