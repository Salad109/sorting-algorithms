package utils

import (
	"fmt"
	"math/rand"
	"strconv"
)

func GenerateFullyRandomArray(arr []int) ([]int, error) {
	fmt.Println("Enter the size of the array:")
	var inputSize string
	_, inputError := fmt.Scanln(&inputSize)
	if inputError != nil {
		return nil, inputError
	}

	size, parseError := strconv.ParseInt(inputSize, 10, 64)
	if parseError != nil {
		return nil, parseError
	}
	if size <= 0 {
		return nil, fmt.Errorf("invalid size: %d", size)
	}
	arr = make([]int, size)
	for i := 0; i < int(size); i++ {
		arr[i] = rand.Intn(10000) // Random number between 0 and 99
	}
	fmt.Println("Array generated successfully.")
	return arr, nil
}
