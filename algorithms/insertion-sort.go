package algorithms

// InsertionSorter implements the Sorter interface using insertion sort algorithm
type InsertionSorter struct{}

// Sort sorts a slice of integers using the insertion sort algorithm
func (is InsertionSorter) Sort(arr []int32) {
	n := len(arr)

	// Start from the second element (index 1)
	for i := 1; i < n; i++ {
		// Store the current element to be inserted
		key := arr[i]

		// Initialize j as the position before current element
		j := i - 1

		// Move elements greater than key one position ahead
		// of their current position
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		// Place the key in its correct position
		arr[j+1] = key
	}
}

// Name returns the name of this sorting algorithm
func (is InsertionSorter) Name() string {
	return "Insertion Sort"
}
