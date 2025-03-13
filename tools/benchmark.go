package tools

import (
	"fmt"
	"sorting-algorithms/algorithms"
	"strconv"
)

// GeneratorInfo pairs a generator function with its descriptive name
type GeneratorInfo struct {
	Func func(int) []int32
	Name string
}

// getDefaultGenerators returns a standard set of array generators
func getDefaultGenerators() []GeneratorInfo {
	return []GeneratorInfo{
		{GenerateFullyRandomArray, "Fully random"},
		{GenerateSortedArray, "Sorted"},
		{GenerateReverseSortedArray, "Reverse sorted"},
		{GenerateOneThirdSortedArray, "33% sorted"},
		{GenerateTwoThirdsSortedArray, "66% sorted"},
	}
}

// getDefaultSizes returns standard array sizes for benchmarking
func getDefaultSizes() []int {
	return []int{100, 2500, 5000, 7500, 10000, 12500, 15000, 17500, 22500, 25000, 30000, 35000, 40000, 45000, 50000, 60000, 70000, 80000}
}

// getDefaultIterations returns the default number of iterations for benchmarking
func getDefaultIterations() int {
	return 100
}

// RunBenchmark performs benchmark tests on the given sorting algorithm
func RunBenchmark(sorter algorithms.Sorter) {
	// Get benchmark configuration
	sizes := getDefaultSizes()
	generators := getDefaultGenerators()
	iterations := getDefaultIterations()

	// Create results matrix [generator][size]
	results := make([][]int64, len(generators))
	for i := range results {
		results[i] = make([]int64, len(sizes))
	}

	// Run benchmarks
	fmt.Printf("\nRunning benchmark for %s with %d iterations per test\n\n",
		sorter.Name(), iterations)

	for genIdx, generator := range generators {
		for sizeIdx, size := range sizes {
			fmt.Printf("Testing with %7d elements, %15s data... ",
				size, generator.Name)

			duration, err := SortArrayIterate(
				size, sorter, generator.Func, iterations, true)

			if err != nil {
				fmt.Printf("\nError: %v\n", err)
				return
			}

			results[genIdx][sizeIdx] = duration.Microseconds()
			fmt.Printf("Average: %v\n", duration)
		}
		fmt.Println()
	}

	// Generate CSV content
	csvData := generateResultCSV(sizes, generators, results)

	// Save results to file
	filename := sorter.Name() + ".csv"
	err := WriteToFile(filename, csvData)
	if err != nil {
		fmt.Printf("Error saving results: %v\n", err)
	} else {
		fmt.Printf("Results saved to %s\n\n", filename)
	}

	// Print summary table
	printResultTable(sizes, generators, results)
}

// generateResultCSV creates CSV content from benchmark results
func generateResultCSV(sizes []int, generators []GeneratorInfo, results [][]int64) []string {
	lines := make([]string, 0, len(sizes)+1)

	// Add header row
	header := "Size"
	for _, gen := range generators {
		header += ";" + gen.Name
	}
	lines = append(lines, header)

	// Add data rows
	for sizeIdx, size := range sizes {
		line := strconv.FormatInt(int64(size), 10)
		for genIdx := range generators {
			line += ";" + strconv.FormatInt(results[genIdx][sizeIdx], 10)
		}
		lines = append(lines, line)
	}

	return lines
}

// printResultTable displays benchmark results in a formatted table
func printResultTable(sizes []int, generators []GeneratorInfo, results [][]int64) {
	fmt.Println("Benchmark results (microseconds):")

	// Print header
	fmt.Printf("%-8s", "Size")
	for _, gen := range generators {
		fmt.Printf(" | %-15s", gen.Name)
	}
	fmt.Println()

	// Print separator
	sep := make([]byte, 8)
	for i := range sep {
		sep[i] = '-'
	}
	fmt.Printf("%-8s", string(sep))

	for range generators {
		fmt.Printf(" | ---------------")
	}
	fmt.Println()

	// Print data rows
	for sizeIdx, size := range sizes {
		fmt.Printf("%-8d", size)
		for genIdx := range generators {
			fmt.Printf(" | %-15d", results[genIdx][sizeIdx])
		}
		fmt.Println()
	}
}
