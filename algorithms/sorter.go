package algorithms

import "golang.org/x/exp/constraints"

// Sorter defines the behavior of a generic sorting algorithm
type Sorter[T constraints.Ordered] interface {
	Sort(arr []T)
	Name() string
}
