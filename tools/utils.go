package tools

import (
	"errors"
	"fmt"
	"time"
)

// PrintArray prints the elements of the array. If the array is empty, it prints a message indicating that.
func PrintArray(arr []int32) {
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
func ValidateSort(arr []int32) bool {
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

func SortArray(size int, sortingAlgorithm func([]int32)) (time.Duration, error) {
	arr := GenerateFullyRandomArray(size)
	// Measure execution time
	start := time.Now()
	sortingAlgorithm(arr)
	elapsed := time.Since(start)

	if !ValidateSort(arr) {
		return elapsed, errors.New("sort unsuccessful")
	}

	return elapsed, nil
}

func SortArrayIterate(size int, sortingAlgorithm func([]int32), iterations int) (time.Duration, error) {
	var totalTime time.Duration

	// Measure average time over multiple runs
	for i := 0; i < iterations; i++ {
		sortingTime, sortingError := SortArray(size, sortingAlgorithm)
		if sortingError != nil {
			return 0, sortingError
		}

		totalTime += sortingTime
	}

	return totalTime / time.Duration(iterations), nil
}
