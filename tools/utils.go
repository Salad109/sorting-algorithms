package tools

import (
	"errors"
	"fmt"
	"math"
	"time"
)

const MaxInt = math.MaxInt32 - 1

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
	fmt.Println("\nIs sorted:", ValidateSort(arr))
}

// ValidateSort checks if the array is sorted in ascending order.
func ValidateSort(arr []int32) bool {
	if len(arr) <= 1 {
		return true
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

// SortArray sorts an array using the provided sorting algorithm and measures the time taken.
func SortArray(size int, sortingAlgorithm func([]int32), generationMethod func(int) []int32, beQuiet bool) (time.Duration, error) {
	arr := generationMethod(size)
	if !beQuiet {
		PrintArray(arr)
	}
	// Measure execution time
	start := time.Now()
	sortingAlgorithm(arr)
	elapsed := time.Since(start)

	if !beQuiet {
		PrintArray(arr)
	}

	if !ValidateSort(arr) {
		return elapsed, errors.New("sort failed validation")
	}

	return elapsed, nil
}

// SortArrayIterate sorts an array multiple times using the provided sorting algorithm and measures the average time taken.
func SortArrayIterate(size int, sortingAlgorithm func([]int32), generationMethod func(int) []int32, iterations int, beQuiet bool) (time.Duration, error) {
	var totalTime time.Duration

	// Measure average time over multiple runs
	for i := 0; i < iterations; i++ {
		if !beQuiet {
			fmt.Println("====== Iteration:", i, "======")
		}
		sortingTime, sortingError := SortArray(size, sortingAlgorithm, generationMethod, beQuiet)
		if sortingError != nil {
			return 0, sortingError
		}
		if !beQuiet {
			fmt.Println("==========================")
		}

		totalTime += sortingTime
	}

	return totalTime / time.Duration(iterations), nil
}
