package main

import (
	"fmt"
	"sorting-algorithms/algorithms"
	"sorting-algorithms/tools"
	"strconv"

	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println("Sorting algorithms benchmark program")
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
	dataType := chooseDataType()
	size := readInt("Enter the size of the array: ", 0, tools.MaxInt)

	if dataType == 1 { // int32
		sorter := chooseSortingAlgorithm[int32]()
		generator := readGenerationMethod(tools.Int32Random)

		fmt.Println("Sorting array using", sorter.Name())
		duration, sortingError := tools.SortArray(size, sorter, generator, false)
		if sortingError != nil {
			fmt.Println("Error sorting array:", sortingError)
			return
		}
		fmt.Println("Array sorted successfully. Time taken:", duration)
	} else { // float32
		sorter := chooseSortingAlgorithm[float32]()
		generator := readGenerationMethod(tools.Float32Random)

		fmt.Println("Sorting array using", sorter.Name())
		duration, sortingError := tools.SortArray(size, sorter, generator, false)
		if sortingError != nil {
			fmt.Println("Error sorting array:", sortingError)
			return
		}
		fmt.Println("Array sorted successfully. Time taken:", duration)
	}
}

// sortArrayMultipleTimes handles the logic for sorting an array multiple times
func sortArrayMultipleTimes() {
	dataType := chooseDataType()
	size := readInt("Enter the size of the array: ", 0, tools.MaxInt)
	iterations := readInt("Enter the number of iterations: ", 1, tools.MaxInt)

	if dataType == 1 { // int32
		sorter := chooseSortingAlgorithm[int32]()
		generator := readGenerationMethod(tools.Int32Random)

		fmt.Println("Sorting array using", sorter.Name())
		averageDuration, sortingError := tools.SortArrayIterate(size, sorter, generator, iterations, false)
		if sortingError != nil {
			fmt.Println("Error sorting array:", sortingError)
			return
		}
		fmt.Println("Array sorted successfully. Average time taken:", averageDuration)
	} else { // float32
		sorter := chooseSortingAlgorithm[float32]()
		generator := readGenerationMethod(tools.Float32Random)

		fmt.Println("Sorting array using", sorter.Name())
		averageDuration, sortingError := tools.SortArrayIterate(size, sorter, generator, iterations, false)
		if sortingError != nil {
			fmt.Println("Error sorting array:", sortingError)
			return
		}
		fmt.Println("Array sorted successfully. Average time taken:", averageDuration)
	}
}

// handleBenchmark delegates to the benchmark package
func handleBenchmark() {
	dataType := chooseDataType()

	if dataType == 1 { // int32
		sorter := chooseSortingAlgorithm[int32]()
		tools.RunBenchmark(sorter, tools.Int32Random, "int32")
	} else { // float32
		sorter := chooseSortingAlgorithm[float32]()
		tools.RunBenchmark(sorter, tools.Float32Random, "float32")
	}
}

// chooseDataType prompts the user to choose between int32 and float32
func chooseDataType() int {
	for {
		fmt.Println("Choose data type:")
		fmt.Println("\t1. Integer (int32)")
		fmt.Println("\t2. Float (float32)")
		choice := readInt("", 1, 2)
		return choice
	}
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

// readGenerationMethod prompts the user to choose a generation method.
func readGenerationMethod[T constraints.Integer | constraints.Float](randomFunc tools.RandomValueGenerator[T]) func(int) []T {
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
			return func(size int) []T { return tools.GenerateFullyRandomArray(size, randomFunc) }
		case 2:
			return tools.GenerateSortedArray[T]
		case 3:
			return tools.GenerateReverseSortedArray[T]
		case 4:
			return func(size int) []T { return tools.GenerateOneThirdSortedArray(size, randomFunc) }
		case 5:
			return func(size int) []T { return tools.GenerateTwoThirdsSortedArray(size, randomFunc) }
		default:
			fmt.Println("Invalid choice. Please try again.")
			continue
		}
	}
}

// chooseSortingAlgorithm prompts the user to choose a sorting algorithm.
func chooseSortingAlgorithm[T constraints.Integer | constraints.Float]() algorithms.Sorter[T] {
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
			return algorithms.BubbleSorter[T]{}
		case 2:
			return algorithms.InsertionSorter[T]{}
		case 3:
			return algorithms.BinaryInsertionSorter[T]{}
		case 4:
			return algorithms.HeapSorter[T]{}
		case 5:
			return algorithms.QuickSorter[T]{}
		default:
			fmt.Println("Invalid choice.")
			continue
		}
	}
}
