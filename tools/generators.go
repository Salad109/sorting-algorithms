package tools

import (
	"math/rand"
	"time"
)

// Global random number generator with proper seeding
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateFullyRandomArrayInt32 generates a fully random array of integers.
func GenerateFullyRandomArrayInt32(size int) []int32 {
	arr := make([]int32, size)
	for i := 0; i < size; i++ {
		arr[i] = int32(rng.Intn(10000)) // Random number between 0 and 9999
	}
	return arr
}

// GenerateFullyRandomArrayFloat32 generates a fully random array of float32 numbers.
func GenerateFullyRandomArrayFloat32(size int) []float32 {
	arr := make([]float32, size)
	for i := 0; i < size; i++ {
		arr[i] = rng.Float32() * 10000 // Random number between 0 and 9999
	}
	return arr
}

// GenerateSortedArrayInt32 generates a sorted array of integers.
func GenerateSortedArrayInt32(size int) []int32 {
	arr := make([]int32, size)
	for i := 0; i < size; i++ {
		arr[i] = int32(i)
	}
	return arr
}

// GenerateSortedArrayFloat32 generates a sorted array of float32 numbers.
func GenerateSortedArrayFloat32(size int) []float32 {
	arr := make([]float32, size)
	for i := 0; i < size; i++ {
		arr[i] = float32(i)
	}
	return arr
}

// GenerateReverseSortedArrayInt32 generates a reverse sorted array of integers.
func GenerateReverseSortedArrayInt32(size int) []int32 {
	arr := make([]int32, size)
	for i := 0; i < size; i++ {
		arr[i] = int32(size - 1 - i)
	}
	return arr
}

// GenerateReverseSortedArrayFloat32 generates a reverse sorted array of float32 numbers.
func GenerateReverseSortedArrayFloat32(size int) []float32 {
	arr := make([]float32, size)
	for i := 0; i < size; i++ {
		arr[i] = float32(size - 1 - i)
	}
	return arr
}

// GenerateOneThirdSortedArrayInt32 generates a 33% sorted array of integers.
func GenerateOneThirdSortedArrayInt32(size int) []int32 {
	arr := make([]int32, size)
	for i := 0; i < size/3; i++ {
		arr[i] = int32(i)
	}
	for i := size / 3; i < size; i++ {
		arr[i] = int32(rng.Intn(10000)) // Random number between 0 and 9999
	}
	return arr
}

// GenerateOneThirdSortedArrayFloat32 generates a 33% sorted array of float32 numbers.
func GenerateOneThirdSortedArrayFloat32(size int) []float32 {
	arr := make([]float32, size)
	for i := 0; i < size/3; i++ {
		arr[i] = float32(i)
	}
	for i := size / 3; i < size; i++ {
		arr[i] = rng.Float32() * 10000 // Random number between 0 and 9999
	}
	return arr
}

// GenerateTwoThirdsSortedArrayInt32 generates a 66% sorted array of integers.
func GenerateTwoThirdsSortedArrayInt32(size int) []int32 {
	arr := make([]int32, size)
	for i := 0; i < (size*2)/3; i++ {
		arr[i] = int32(i)
	}
	for i := (size * 2) / 3; i < size; i++ {
		arr[i] = int32(rng.Intn(10000)) // Random number between 0 and 9999
	}
	return arr
}

// GenerateTwoThirdsSortedArrayFloat32 generates a 66% sorted array of float32 numbers.
func GenerateTwoThirdsSortedArrayFloat32(size int) []float32 {
	arr := make([]float32, size)
	for i := 0; i < (size*2)/3; i++ {
		arr[i] = float32(i)
	}
	for i := (size * 2) / 3; i < size; i++ {
		arr[i] = rng.Float32() * 10000 // Random number between 0 and 9999
	}
	return arr
}
