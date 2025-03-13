package tools

import (
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"sorting-algorithms/algorithms"
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
func SortArray(size int, sorter algorithms.Sorter, generationMethod func(int) []int32, beQuiet bool) (time.Duration, error) {
	arr := generationMethod(size)
	if !beQuiet {
		PrintArray(arr)
	}
	// Measure execution time
	start := time.Now()
	sorter.Sort(arr)
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
func SortArrayIterate(size int, sorter algorithms.Sorter, generationMethod func(int) []int32, iterations int, beQuiet bool) (time.Duration, error) {
	times := make([]time.Duration, iterations)

	// Measure average time over multiple runs
	for i := 0; i < iterations; i++ {
		if !beQuiet {
			fmt.Println("====== Iteration:", i, "======")
		}
		sortingTime, sortingError := SortArray(size, sorter, generationMethod, beQuiet)
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

// WriteToFile writes the benchmark results to a file.
func WriteToFile(filename string, results []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	for _, line := range results {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}

	fmt.Println("Results written to", filename)
	return nil
}

func TrimExtremes(arr []time.Duration) []time.Duration {
	if len(arr) == 0 {
		return []time.Duration{}
	}
	if len(arr) <= 2 {
		return arr
	}
	// Sort the array
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	// Calculate the number of elements to remove
	numToRemove := int(float64(len(arr)) * 0.2)
	if numToRemove < 1 {
		numToRemove = 1
	}
	// Remove the smallest and largest elements
	trimmedArr := arr[numToRemove : len(arr)-numToRemove]
	return trimmedArr
}
