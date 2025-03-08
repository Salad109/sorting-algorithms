package utils

import "fmt"

func PrintArray(arr []int) {
	if len(arr) == 0 {
		fmt.Println("Array is empty.")
		return
	}
	if len(arr) > 20 {
		for i := 0; i < 20; i++ {
			fmt.Print(arr[i], " ")
		}
		fmt.Println("...")
		fmt.Println("Is sorted:", ValidateSort(arr))
		return
	}
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println("Is sorted:", ValidateSort(arr))

	fmt.Println()
	return
}

func ValidateSort(arr []int) bool {
	if len(arr) <= 1 {
		return false
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
