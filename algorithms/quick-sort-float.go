package algorithms

// QuickFloatSorter implements the FloatSorter interface using quick sort algorithm
type QuickFloatSorter struct{}

// SortFloat sorts a slice of floats using the quick sort algorithm
func (qfs QuickFloatSorter) SortFloat(arr []float32) {
	quickSortFloat(arr, 0, len(arr)-1)
}

// quickSort is a recursive function that implements the quick sort algorithm
func quickSortFloat(arr []float32, low int, high int) {
	if low < high {
		// Partition the array
		pi := partitionFloat(arr, low, high)

		// Recursively sort elements before and after partition
		quickSortFloat(arr, low, pi-1)
		quickSortFloat(arr, pi+1, high)
	}
}

// partition rearranges the elements in the array and returns the index of the pivot
func partitionFloat(arr []float32, low int, high int) int {
	pivot := arr[high] // pivot
	i := low - 1       // Index of smaller element

	for j := low; j < high; j++ {
		// If current element is smaller than or equal to pivot
		if arr[j] <= pivot {
			i++                             // increment index of smaller element
			arr[i], arr[j] = arr[j], arr[i] // Swap
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] // Swap pivot to the correct position
	return i + 1
}

// Name returns the name of this sorting algorithm
func (qfs QuickFloatSorter) Name() string {
	return "Quick Float Sort"
}
