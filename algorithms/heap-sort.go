package algorithms

import "golang.org/x/exp/constraints"

// HeapSorter implements the Sorter interface using heap sort algorithm
type HeapSorter[T constraints.Ordered] struct{}

// Sort sorts a slice using the heap sort algorithm
func (hs HeapSorter[T]) Sort(arr []T) {
	n := len(arr)
	// Build a max heap
	for i := n/2 - 1; i >= 0; i-- {
		hs.heapify(arr, n, i)
	}
	// One by one extract elements from heap
	for i := n - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i] // Swap
		hs.heapify(arr, i, 0)
	}
}

// heapify function to maintain the heap property
func (hs HeapSorter[T]) heapify(arr []T, n int, i int) {
	largest := i     // Initialize largest as root
	left := 2*i + 1  // left child index
	right := 2*i + 2 // right child index

	// If left child is larger than root
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// If right child is larger than largest so far
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i] // Swap
		hs.heapify(arr, n, largest)
	}
}

// Name returns the name of this sorting algorithm
func (hs HeapSorter[T]) Name() string {
	return "Heap Sort"
}
