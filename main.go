package main

import (
	"fmt"
	"sorting-algorithms/algorithms"
	"sorting-algorithms/tools"
	"strconv"
)

func main() {
	var arr []int32
	fmt.Println("Sorting algorithms benchmark program")
	action := ""
	for action != "0" {
		fmt.Println("Choose an action:")
		fmt.Println("\t0. Exit")
		fmt.Println("\t1. Generate a fully random array")
		fmt.Println("\t2. Print array")
		fmt.Println("\t3. Sort array using bubble sort")
		fmt.Println("\t4. Sort array using bubble sort multiple times")

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

			arr = tools.GenerateFullyRandomArray(int(size))
			fmt.Println("Array generated successfully.")
		case "2":
			tools.PrintArray(arr)
		case "3":
			fmt.Println("Sorting array using Bubble Sort...")
			duration, sortingError := tools.SortArray(arr, algorithms.BubbleSort)
			if sortingError != nil {
				fmt.Println("Error sorting array:", sortingError)
				continue
			}
			fmt.Println("Array sorted successfully. Time taken:", duration)
		case "4":
			fmt.Println("Sorting array using Bubble Sort multiple times...")
			duration, sortingError := tools.SortArrayIterate(arr, algorithms.BubbleSort, 10)
			if sortingError != nil {
				fmt.Println("Error sorting array:", sortingError)
				continue
			}
			fmt.Println("Array sorted successfully. Average time taken:", duration)
		default:
			fmt.Println("Invalid action. Please try again.")
		}
	}
}
