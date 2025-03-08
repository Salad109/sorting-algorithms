package algorithms

import "errors"

// BubbleSort sorts a slice of integers using the bubble sort algorithm.
func BubbleSort(arr []int) error {
	n := len(arr)
	if n <= 1 {
		return errors.New("array too small to sort")
	}

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
	return nil
}
