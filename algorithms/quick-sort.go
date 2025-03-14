package algorithms

import "golang.org/x/exp/constraints"

// QuickSorter implements the Sorter interface using quick sort algorithm
type QuickSorter[T constraints.Ordered] struct{}

// Sort sorts a slice using the quick sort algorithm
func (qs QuickSorter[T]) Sort(arr []T) {
	quickSort(arr, 0, len(arr)-1)
}

// quickSort is a recursive function that implements the quick sort algorithm
func quickSort[T constraints.Ordered](arr []T, low int, high int) {
	if low < high {
		// Partition the array
		pi := partition[T](arr, low, high)

		// Recursively sort elements before and after partition
		quickSort[T](arr, low, pi-1)
		quickSort[T](arr, pi+1, high)
	}
}

// partition rearranges the elements in the array and returns the index of the pivot
func partition[T constraints.Ordered](arr []T, low int, high int) int {
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
func (qs QuickSorter[T]) Name() string {
	return "Quick Sort"
}
