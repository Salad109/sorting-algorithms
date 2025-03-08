package tools

import (
	"math/rand"
)

// GenerateFullyRandomArray generates a fully random array of integers.
func GenerateFullyRandomArray(size int64) []int {

	arr := make([]int, size)
	var i int64
	for i = 0; i < size; i++ {
		arr[i] = rand.Intn(10000) // Random number between 0 and 9999
	}
	return arr
}
