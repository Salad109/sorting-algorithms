package tools

import (
	"fmt"
	"sorting-algorithms/algorithms"
	"strconv"
)

// RunBenchmarkFloat32 performs benchmark tests on the given float sorting algorithm
func RunBenchmarkFloat32(sorter algorithms.FloatSorter) {
	// Get benchmark configuration
	sizes := getDefaultSizes()
	generators := getDefaultGeneratorsFloat32()
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

			duration, err := SortFloatArrayIterate(
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
	csvData := generateResultCSVFloat32(sizes, generators, results)

	// Save results to file
	filename := sorter.Name() + " float32.csv"
	err := WriteToFile(filename, csvData)
	if err != nil {
		fmt.Printf("Error saving results: %v\n", err)
	} else {
		fmt.Printf("Results saved to %s\n\n", filename)
	}

	// Print summary table
	printResultTableFloat32(sizes, generators, results)
}

// generateResultCSVFloat32 creates CSV content from benchmark results
func generateResultCSVFloat32(sizes []int, generators []GeneratorInfoFloat32, results [][]int64) []string {
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

// printResultTableFloat32 displays benchmark results in a formatted table
func printResultTableFloat32(sizes []int, generators []GeneratorInfoFloat32, results [][]int64) {
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
