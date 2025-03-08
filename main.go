package main

import (
	"fmt"
	"sorting-algorithms/algorithms"
	"sorting-algorithms/tools"
	"strconv"
)

func main() {
	fmt.Println("Sorting algorithms benchmark program")
	action := ""
	for action != "0" {
		fmt.Println("Choose an action:")
		fmt.Println("\t0. Exit")
		fmt.Println("\t1. Sort array using bubble sort")
		fmt.Println("\t2. Sort array using bubble sort multiple times")

		_, inputError := fmt.Scanln(&action)
		if inputError != nil {
			fmt.Println("Error reading action input:", inputError)
			return
		}

		switch action {
		case "0":
			return
		case "1":
			fmt.Println("Enter the size of the array:")
			var inputSize string
			_, inputError := fmt.Scanln(&inputSize)
			if inputError != nil {
				fmt.Println("Error reading size input:", inputError)
				return
			}

			size, parseError := strconv.ParseInt(inputSize, 10, 32)
			if parseError != nil {
				fmt.Println("Error parsing size:", parseError)
				return
			}

			if size <= 0 {
				fmt.Println("Size must be greater than 0.")
				continue
			}

			fmt.Println("Sorting array using Bubble Sort...")
			duration, sortingError := tools.SortArray(int(size), algorithms.BubbleSort)
			if sortingError != nil {
				fmt.Println("Error sorting array:", sortingError)
				continue
			}
			fmt.Println("Array sorted successfully. Time taken:", duration)
		case "2":
			fmt.Println("Enter the size of the array:")
			var inputSize string
			_, inputSizeError := fmt.Scanln(&inputSize)
			if inputSizeError != nil {
				fmt.Println("Error reading size input:", inputSizeError)
				return
			}

			size, parseError := strconv.ParseInt(inputSize, 10, 32)
			if parseError != nil {
				fmt.Println("Error parsing size:", parseError)
				return
			}

			if size <= 0 {
				fmt.Println("Size must be greater than 0.")
				continue
			}

			fmt.Println("Enter the iteration count:")
			var inputIterations string
			_, inputIterationError := fmt.Scanln(&inputIterations)
			if inputIterationError != nil {
				fmt.Println("Error reading iteration count input:", inputIterationError)
				return
			}
			iterations, parseError := strconv.ParseInt(inputIterations, 10, 64)
			if parseError != nil {
				fmt.Println("Error parsing iteration count:", parseError)
				continue
			}
			if iterations <= 0 {
				fmt.Println("Iteration count must be greater than 0.")
				continue
			}
			averageDuration, sortingError := tools.SortArrayIterate(int(size), algorithms.BubbleSort, int(iterations))
			if sortingError != nil {
				fmt.Println("Error sorting array:", sortingError)
				continue
			}
			fmt.Println("Array sorted successfully. Average time taken:", averageDuration)
		default:
			fmt.Println("Invalid action. Please try again.")
		}
	}
}
