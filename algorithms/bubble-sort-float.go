package algorithms

// BubbleFloatSorter implements the FloatSorter interface using bubble sort algorithm
type BubbleFloatSorter struct{}

// SortFloat sorts a slice of floats using the bubble sort algorithm
func (bfs BubbleFloatSorter) SortFloat(arr []float32) {
	n := len(arr)

	for {
		swapped := false
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i] // Swap elements
				swapped = true
			}
		}
		if !swapped {
			break
		}
		n-- // Reduce the range of the next pass
	}
}

// Name returns the name of this sorting algorithm
func (bfs BubbleFloatSorter) Name() string {
	return "Bubble Float Sort"
}
