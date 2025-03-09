package main

import (
	"fmt"
	"sorting-algorithms/algorithms"
	"sorting-algorithms/tools"
	"strconv"
)

func main() {
	fmt.Println("Sorting algorithms runBenchmark program")
	for {
		fmt.Println("Choose an action:")
		fmt.Println("\t0. Exit")
		fmt.Println("\t1. Sort array once")
		fmt.Println("\t2. Sort array multiple times")
		fmt.Println("\t3. Benchmark algorithm")

		action := readInt("", 0, 3)

		switch action {
		case 0:
			return
		case 1:
			sortArrayOnce()
		case 2:
			sortArrayMultipleTimes()
		case 3:
			handleBenchmark()
		default:
			fmt.Println("Invalid action. Please try again.")
		}
	}
}

// sortArrayOnce handles the logic for sorting an array once
func sortArrayOnce() {
	algorithm, name := chooseSortingAlgorithm()
	size := readInt("Enter the size of the array: ", 0, tools.MaxInt)
	generator := readGenerationMethod()

	fmt.Println("Sorting array using", name)
	duration, sortingError := tools.SortArray(size, algorithm, generator, false)
	if sortingError != nil {
		fmt.Println("Error sorting array:", sortingError)
		return
	}
	fmt.Println("Array sorted successfully. Time taken:", duration)
}

// sortArrayMultipleTimes handles the logic for sorting an array multiple times
func sortArrayMultipleTimes() {
	algorithm, name := chooseSortingAlgorithm()
	size := readInt("Enter the size of the array: ", 0, tools.MaxInt)
	generator := readGenerationMethod()
	iterations := readInt("Enter the number of iterations: ", 1, tools.MaxInt)

	fmt.Println("Sorting array using", name)
	averageDuration, sortingError := tools.SortArrayIterate(size, algorithm, generator, iterations, false)
	if sortingError != nil {
		fmt.Println("Error sorting array:", sortingError)
		return
	}
	fmt.Println("Array sorted successfully. Average time taken:", averageDuration)
}

// runBenchmark runs a runBenchmark for the given sorting algorithm
func runBenchmark(sortingAlgorithm func([]int32), sortingAlgorithmName string) {
	generationMethodList := []struct {
		generator func(int) []int32
		name      string
	}{
		{tools.GenerateFullyRandomArray, "Fully random"},
		{tools.GenerateSortedArray, "Sorted"},
		{tools.GenerateReverseSortedArray, "Reverse sorted"},
		{tools.GenerateOneThirdSortedArray, "33 sorted"},
		{tools.GenerateTwoThirdsSortedArray, "66 sorted"},
	}
	sizes := []int{100, 2000, 4000, 6000, 8000, 10000, 12000, 14000, 16000, 18000, 20000}
	results := make([][]int64, len(generationMethodList))
	iterations := 100
	for generatorIndex, generator := range generationMethodList {
		for _, size := range sizes {
			fmt.Print("Benchmarking ", sortingAlgorithmName, " Size: ", size, " Generation method: ", generator.name, "... ")
			averageDuration, sortingError := tools.SortArrayIterate(size, sortingAlgorithm, generator.generator, iterations, true)
			if sortingError != nil {
				fmt.Println("Error sorting array:", sortingError)
				return
			}
			results[generatorIndex] = append(results[generatorIndex], averageDuration.Microseconds())
			fmt.Println("Average time taken:", averageDuration)
		}
	}
	fileContent := make([]string, 0)
	header := "Size"
	for _, generator := range generationMethodList {
		header += ";" + generator.name
	}
	fileContent = append(fileContent, header)
	for sizeIndex, size := range sizes {
		line := strconv.FormatInt(int64(size), 10)
		for generatorIndex := 0; generatorIndex < len(generationMethodList); generatorIndex++ {
			line += ";" + strconv.FormatInt(results[generatorIndex][sizeIndex], 10)
		}
		fileContent = append(fileContent, line)
	}
	fmt.Println("Benchmark results:")
	for _, line := range fileContent {
		fmt.Println(line)
	}
	tools.WriteToFile(sortingAlgorithmName+".csv", fileContent)
	return
}

func handleBenchmark() {
	algorithm, name := chooseSortingAlgorithm()
	runBenchmark(algorithm, name)
}

// readInt reads an integer from the user with a prompt and validates it against min and max values.
func readInt(prompt string, min, max int) int {
	for {
		fmt.Print(prompt)
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		val, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid number, please try again.")
			continue
		}

		if val < min || val > max {
			fmt.Printf("Value must be between %d and %d, please try again.\n", min, max)
			continue
		}

		return val
	}
}

// readGenerationMethod prompts the user to choose a generation method for the array.
func readGenerationMethod() func(int) []int32 {
	for {
		fmt.Println("Choose a generation method:")
		fmt.Println("\t1. Fully random")
		fmt.Println("\t2. Sorted")
		fmt.Println("\t3. Reverse sorted")
		fmt.Println("\t4. 33% sorted")
		fmt.Println("\t5. 66% sorted")
		choice := readInt("", 1, 5)
		switch choice {
		case 1:
			return tools.GenerateFullyRandomArray
		case 2:
			return tools.GenerateSortedArray
		case 3:
			return tools.GenerateReverseSortedArray
		case 4:
			return tools.GenerateOneThirdSortedArray
		case 5:
			return tools.GenerateTwoThirdsSortedArray
		default:
			fmt.Println("Invalid choice. Please try again.")
			continue
		}
	}
}

// chooseSortingAlgorithm prompts the user to choose a sorting algorithm.
func chooseSortingAlgorithm() (func([]int32), string) {
	for {
		fmt.Println("Choose a sorting algorithm:")
		fmt.Println("\t1. Bubble Sort")
		fmt.Println("\t2. Insertion Sort")
		fmt.Println("\t3. Binary Insertion Sort")
		fmt.Println("\t4. Heap Sort")
		fmt.Println("\t5. Quick Sort")
		choice := readInt("", 1, 5)
		switch choice {
		case 1:
			return algorithms.BubbleSort, "Bubble Sort"
		case 2:
			return algorithms.InsertionSort, "Insertion Sort"
		case 3:
			return algorithms.BinaryInsertionSort, "Binary Insertion Sort"
		case 4:
			return algorithms.HeapSort, "Heap Sort"
		case 5:
			return algorithms.QuickSort, "Quick Sort"
		default:
			fmt.Println("Invalid choice.")
			continue
		}
	}
}
