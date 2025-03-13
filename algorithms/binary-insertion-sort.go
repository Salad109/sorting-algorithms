package algorithms

// BinaryInsertionSorter implements the Sorter interface using binary insertion sort
type BinaryInsertionSorter struct{}

// Sort sorts a slice of integers using the binary insertion sort algorithm
func (bis BinaryInsertionSorter) Sort(arr []int32) {
	n := len(arr)

	for i := 1; i < n; i++ {
		// Store the current element to be inserted
		key := arr[i]

		// Find the position where key should be inserted using binary search
		insertionPoint := bis.binarySearch(arr, 0, i-1, key)

		// Shift all elements to the right to make space for the key
		if insertionPoint < i {
			// Save the element at position i
			temp := arr[i]

			// Shift all elements between insertionPoint and i-1
			for j := i; j > insertionPoint; j-- {
				arr[j] = arr[j-1]
			}

			// Insert the key at the correct position
			arr[insertionPoint] = temp
		}
	}
}

// binarySearch finds the position where x should be inserted in a sorted array
// such that all elements before this position are less than or equal to x
// and all elements after this position are greater than x
func (bis BinaryInsertionSorter) binarySearch(arr []int32, left, right int, x int32) int {
	if right <= left {
		if arr[left] > x {
			return left
		}
		return left + 1
	}

	mid := (left + right) / 2

	if x == arr[mid] {
		return mid + 1
	}

	if x > arr[mid] {
		return bis.binarySearch(arr, mid+1, right, x)
	}

	return bis.binarySearch(arr, left, mid-1, x)
}

// Name returns the name of this sorting algorithm
func (bis BinaryInsertionSorter) Name() string {
	return "Binary Insertion Sort"
}
