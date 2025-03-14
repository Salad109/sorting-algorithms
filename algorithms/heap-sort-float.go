package algorithms

// HeapFloatSorter implements the FloatSorter interface using heap sort algorithm
type HeapFloatSorter struct{}

// SortFloat sorts a slice of floats using the heap sort algorithm
func (hfs HeapFloatSorter) SortFloat(arr []float32) {
	n := len(arr)
	// Build a max heap
	for i := n/2 - 1; i >= 0; i-- {
		hfs.heapify(arr, n, i)
	}
	// One by one extract elements from heap
	for i := n - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i] // Swap
		hfs.heapify(arr, i, 0)
	}
}

// heapify function to maintain the heap property
func (hfs HeapFloatSorter) heapify(arr []float32, n int, i int) {
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
		hfs.heapify(arr, n, largest)
	}
}

// Name returns the name of this sorting algorithm
func (hfs HeapFloatSorter) Name() string {
	return "Heap Float Sort"
}
