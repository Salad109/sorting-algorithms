package algorithms

// Sorter defines the behavior of a sorting algorithm
type Sorter interface {
	Sort(arr []int32)
	Name() string
}
