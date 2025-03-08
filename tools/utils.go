package tools

import (
	"errors"
	"fmt"
	"time"
)

// PrintArray prints the elements of the array. If the array is empty, it prints a message indicating that.
func PrintArray(arr []int) {
	if len(arr) == 0 {
		fmt.Println("Array is empty.")
		return
	}
	if len(arr) > 20 {
		for i := 0; i < 20; i++ {
			fmt.Print(arr[i], " ")
		}
		fmt.Println("...")
		fmt.Println("Is sorted:", ValidateSort(arr))
		return
	}
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println("Is sorted:", ValidateSort(arr))

	fmt.Println()
	return
}

// ValidateSort checks if the array is sorted in ascending order.
func ValidateSort(arr []int) bool {
	if len(arr) <= 1 {
		return false
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

// SortArray sorts the given array using the provided sorting algorithm and returns the time taken to sort.
func SortArray(arr []int, sortingAlgorithm func([]int)) (time.Duration, error) {
	if len(arr) <= 1 {
		return 0, errors.New("array too small to sort")
	}

	// Measure execution time
	start := time.Now()
	sortingAlgorithm(arr)
	elapsed := time.Since(start)

	return elapsed, nil
}

// SortArrayIterate sorts the given array using the provided sorting algorithm multiple times and returns the average time taken to sort.
func SortArrayIterate(arr []int, sortingAlgorithm func([]int), iterations int) (time.Duration, error) {
	var totalTime time.Duration

	// Discard first run (warmup)
	tempArr := make([]int, len(arr))
	copy(tempArr, arr)
	_, err := SortArray(tempArr, sortingAlgorithm)
	if err != nil {
		return 0, err
	}

	// Measure average time over multiple runs
	for i := 0; i < iterations; i++ {
		tempArr := make([]int, len(arr))
		copy(tempArr, arr)

		start := time.Now()
		sortingAlgorithm(tempArr)
		totalTime += time.Since(start)
	}

	return totalTime / time.Duration(iterations), nil
}
