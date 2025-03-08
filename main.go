package main

import (
	"fmt"
	"sorting-algorithms/algorithms"
	"sorting-algorithms/utils"
)

func main() {
	var arr []int
	fmt.Println("Sorting algorithms benchmark program")
	action := ""
	for action != "0" {
		fmt.Println("Choose an action:")
		fmt.Println("\t0. Exit")
		fmt.Println("\t1. Generate Array")
		fmt.Println("\t2. Print Array")
		fmt.Println("\t3. Sort Array using Bubble Sort")

		_, inputError := fmt.Scanln(&action)
		if inputError != nil {
			fmt.Println("Error reading action input:", inputError)
			return
		}

		var err error
		switch action {
		case "0":
			return
		case "1":
			arr, err = utils.GenerateFullyRandomArray(arr)
			if err != nil {
				fmt.Println("Error generating array:", err)
				continue
			}
		case "2":
			utils.PrintArray(arr)
		case "3":
			fmt.Println("Sorting array using Bubble Sort...")
			sortingError := algorithms.BubbleSort(arr)
			if sortingError != nil {
				fmt.Println("Error sorting array:", sortingError)
				continue
			}
			fmt.Println("Array sorted successfully.")
		}
	}
}
