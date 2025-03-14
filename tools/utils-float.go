package tools

import (
	"errors"
	"fmt"
	"sorting-algorithms/algorithms"
	"time"
)

// PrintFloatArray prints the elements of the float array. If the array is empty, it prints a message indicating that.
func PrintFloatArray(arr []float32) {
	if len(arr) == 0 {
		fmt.Println("Array is empty.")
		return
	}
	if len(arr) > 20 {
		for i := 0; i < 20; i++ {
			fmt.Printf("%.2f ", arr[i])
		}
		fmt.Println("...")
		fmt.Println("Is sorted:", ValidateFloatSort(arr))
		return
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%.2f ", arr[i])
	}
	fmt.Println("\nIs sorted:", ValidateFloatSort(arr))
}

// ValidateFloatSort checks if the float array is sorted in ascending order.
func ValidateFloatSort(arr []float32) bool {
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

// SortFloatArray sorts a float array using the provided sorting algorithm and measures the time taken.
func SortFloatArray(size int, sorter algorithms.FloatSorter, generationMethod func(int) []float32, beQuiet bool) (time.Duration, error) {
	arr := generationMethod(size)
	if !beQuiet {
		PrintFloatArray(arr)
	}
	// Measure execution time
	start := time.Now()
	sorter.SortFloat(arr)
	elapsed := time.Since(start)

	if !beQuiet {
		PrintFloatArray(arr)
	}

	if !ValidateFloatSort(arr) {
		return elapsed, errors.New("sort failed validation")
	}

	return elapsed, nil
}

// SortFloatArrayIterate sorts a float array multiple times using the provided sorting algorithm and measures the average time taken.
func SortFloatArrayIterate(size int, sorter algorithms.FloatSorter, generationMethod func(int) []float32, iterations int, beQuiet bool) (time.Duration, error) {
	times := make([]time.Duration, iterations)

	// Measure average time over multiple runs
	for i := 0; i < iterations; i++ {
		if !beQuiet {
			fmt.Println("====== Iteration:", i, "======")
		}
		sortingTime, sortingError := SortFloatArray(size, sorter, generationMethod, beQuiet)
		if sortingError != nil {
			return 0, sortingError
		}
		if !beQuiet {
			fmt.Println("==========================")
		}

		times[i] = sortingTime
	}

	// Trim the extremes
	times = TrimExtremes(times)
	if len(times) == 0 {
		return 0, errors.New("no valid times to calculate average")
	}
	// Calculate average time
	totalTime := time.Duration(0)
	for _, t := range times {
		totalTime += t
	}
	averageTime := totalTime / time.Duration(len(times))
	if !beQuiet {
		fmt.Println("Average time taken:", averageTime)
	}
	return averageTime, nil
}
