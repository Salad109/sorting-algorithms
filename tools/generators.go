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
