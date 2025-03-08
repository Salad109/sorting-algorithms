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
			size, err := readInt("Enter the size of the array: ", 1)
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}

			fmt.Println("Sorting array using Bubble Sort...")
			duration, sortingError := tools.SortArray(size, algorithms.BubbleSort)
			if sortingError != nil {
				fmt.Println("Error sorting array:", sortingError)
				continue
			}
			fmt.Println("Array sorted successfully. Time taken:", duration)
		case "2":
			size, err := readInt("Enter the size of the array: ", 1)
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}

			iterations, err2 := readInt("Enter the number of iterations: ", 1)
			if err2 != nil {
				fmt.Println("Error reading input:", err2)
				return
			}

			averageDuration, sortingError := tools.SortArrayIterate(size, algorithms.BubbleSort, iterations)
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

func readInt(prompt string, min int) (int, error) {
	for {
		fmt.Print(prompt)
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			return 0, err
		}

		val, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid number, please try again.")
			continue
		}

		if val < min {
			fmt.Printf("Value must be at least %d, please try again.\n", min)
			continue
		}

		return val, nil
	}
}
