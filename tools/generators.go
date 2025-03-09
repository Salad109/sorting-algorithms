package tools

import (
	"math/rand"
)

// GenerateFullyRandomArray generates a fully random array of integers.
func GenerateFullyRandomArray(size int) []int32 {

	arr := make([]int32, size)
	for i := 0; i < size; i++ {
		arr[i] = int32(rand.Intn(10000)) // Random number between 0 and 9999
	}
	return arr
}

// GenerateSortedArray generates a sorted array of integers.
func GenerateSortedArray(size int) []int32 {

	arr := make([]int32, size)
	for i := 0; i < size; i++ {
		arr[i] = int32(i)
	}
	return arr
}

// GenerateReverseSortedArray generates a reverse sorted array of integers.
func GenerateReverseSortedArray(size int) []int32 {
	arr := make([]int32, size)
	for i := 0; i < size; i++ {
		arr[i] = int32(size - 1 - i)
	}
	return arr
}

// GenerateOneThirdSortedArray generates a 33% sorted array of integers.
func GenerateOneThirdSortedArray(size int) []int32 {
	arr := make([]int32, size)
	for i := 0; i < size/3; i++ {
		arr[i] = int32(i)
	}
	for i := size / 3; i < size; i++ {
		arr[i] = int32(rand.Intn(10000)) // Random number between 0 and 9999
	}
	return arr
}

// GenerateTwoThirdsSortedArray generates a 66% sorted array of integers.
func GenerateTwoThirdsSortedArray(size int) []int32 {
	arr := make([]int32, size)
	for i := 0; i < (size*2)/3; i++ {
		arr[i] = int32(i)
	}
	for i := (size * 2) / 3; i < size; i++ {
		arr[i] = int32(rand.Intn(10000)) // Random number between 0 and 9999
	}
	return arr
}
