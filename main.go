package main

import (
	"fmt"
	"sorting-algorithms/algorithms"
	"sorting-algorithms/tools"
	"strconv"
)

func main() {
	fmt.Println("Sorting algorithms benchmark program")
	for {
		fmt.Println("Choose an action:")
		fmt.Println("\t0. Exit")
		fmt.Println("\t1. Sort array once")
		fmt.Println("\t2. Sort array multiple times")

		action := readInt("", 0, 2)

		switch action {
		case 0:
			return
		case 1:
			sortArrayOnce()
		case 2:
			sortArrayMultipleTimes()
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
	duration, sortingError := tools.SortArray(size, algorithm, generator)
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
	averageDuration, sortingError := tools.SortArrayIterate(size, algorithm, generator, iterations)
	if sortingError != nil {
		fmt.Println("Error sorting array:", sortingError)
		return
	}
	fmt.Println("Array sorted successfully. Average time taken:", averageDuration)
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
		choice := readInt("", 1, 3)
		switch choice {
		case 1:
			return algorithms.BubbleSort, "Bubble Sort"
		case 2:
			return algorithms.InsertionSort, "Insertion Sort"
		case 3:
			return algorithms.BinaryInsertionSort, "Binary Insertion Sort"
		default:
			fmt.Println("Invalid choice.")
			continue
		}
	}
}
