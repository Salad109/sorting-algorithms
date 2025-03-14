package tools

import (
	"fmt"
	"sorting-algorithms/algorithms"
	"strconv"
)

// GeneratorInfoInt32 pairs a generator function with its descriptive name
type GeneratorInfoInt32 struct {
	Func func(int) []int32
	Name string
}

// GeneratorInfoFloat32 pairs a generator function with its descriptive name
type GeneratorInfoFloat32 struct {
	Func func(int) []float32
	Name string
}

// getDefaultGeneratorsInt32 returns a standard set of array generators
func getDefaultGeneratorsInt32() []GeneratorInfoInt32 {
	return []GeneratorInfoInt32{
		{GenerateFullyRandomArrayInt32, "Fully random"},
		{GenerateSortedArrayInt32, "Sorted"},
		{GenerateReverseSortedArrayInt32, "Reverse sorted"},
		{GenerateOneThirdSortedArrayInt32, "33% sorted"},
		{GenerateTwoThirdsSortedArrayInt32, "66% sorted"},
	}
}

// getDefaultGeneratorsFloat32 returns a standard set of array generators for float32
func getDefaultGeneratorsFloat32() []GeneratorInfoFloat32 {
	return []GeneratorInfoFloat32{
		{GenerateFullyRandomArrayFloat32, "Fully random"},
		{GenerateSortedArrayFloat32, "Sorted"},
		{GenerateReverseSortedArrayFloat32, "Reverse sorted"},
		{GenerateOneThirdSortedArrayFloat32, "33% sorted"},
		{GenerateTwoThirdsSortedArrayFloat32, "66% sorted"},
	}
}

// getDefaultSizes returns standard array sizes for benchmarking
func getDefaultSizes() []int {
	return []int{100, 2500, 5000, 7500, 10000, 12500, 15000, 17500, 20000, 22500, 25000, 30000, 35000, 40000, 45000, 50000, 60000, 70000, 80000}
}

// getDefaultIterations returns the default number of iterations for benchmarking
func getDefaultIterations() int {
	return 100
}

// RunBenchmarkInt32 performs benchmark tests on the given sorting algorithm
func RunBenchmarkInt32(sorter algorithms.Sorter) {
	// Get benchmark configuration
	sizes := getDefaultSizes()
	generators := getDefaultGeneratorsInt32()
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
	csvData := generateResultCSVInt32(sizes, generators, results)

	// Save results to file
	filename := sorter.Name() + " int32.csv"
	err := WriteToFile(filename, csvData)
	if err != nil {
		fmt.Printf("Error saving results: %v\n", err)
	} else {
		fmt.Printf("Results saved to %s\n\n", filename)
	}

	// Print summary table
	printResultTableInt32(sizes, generators, results)
}

// generateResultCSVInt32 creates CSV content from benchmark results
func generateResultCSVInt32(sizes []int, generators []GeneratorInfoInt32, results [][]int64) []string {
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

// printResultTableInt32 displays benchmark results in a formatted table
func printResultTableInt32(sizes []int, generators []GeneratorInfoInt32, results [][]int64) {
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
