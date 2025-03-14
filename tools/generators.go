package tools

import (
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

// Global random number generator with proper seeding
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// RandomValueGenerator represents a function that generates random values of type T
type RandomValueGenerator[T any] func() T

// Int32Random generates a random int32 value
func Int32Random() int32 {
	return int32(rng.Intn(10000)) // Random number between 0 and 9999
}

// Float32Random generates a random float32 value
func Float32Random() float32 {
	return rng.Float32() * 10000 // Random number between 0 and 9999
}

// GenerateFullyRandomArray generates a fully random array
func GenerateFullyRandomArray[T any](size int, randomFunc RandomValueGenerator[T]) []T {
	arr := make([]T, size)
	for i := 0; i < size; i++ {
		arr[i] = randomFunc()
	}
	return arr
}

// GenerateSortedArray generates a sorted array
func GenerateSortedArray[T constraints.Integer | constraints.Float](size int) []T {
	arr := make([]T, size)
	for i := 0; i < size; i++ {
		arr[i] = T(i)
	}
	return arr
}

// GenerateReverseSortedArray generates a reverse sorted array
func GenerateReverseSortedArray[T constraints.Integer | constraints.Float](size int) []T {
	arr := make([]T, size)
	for i := 0; i < size; i++ {
		arr[i] = T(size - 1 - i)
	}
	return arr
}

// GenerateOneThirdSortedArray generates a 33% sorted array
func GenerateOneThirdSortedArray[T constraints.Integer | constraints.Float](size int, randomFunc RandomValueGenerator[T]) []T {
	arr := make([]T, size)

	// Fill the first third with sequential values
	for i := 0; i < size/3; i++ {
		arr[i] = T(i)
	}

	// Fill the rest with random values
	for i := size / 3; i < size; i++ {
		arr[i] = randomFunc()
	}
	return arr
}

// GenerateTwoThirdsSortedArray generates a 66% sorted array
func GenerateTwoThirdsSortedArray[T constraints.Integer | constraints.Float](size int, randomFunc RandomValueGenerator[T]) []T {
	arr := make([]T, size)

	// Fill the first two thirds with sequential values
	for i := 0; i < (size*2)/3; i++ {
		arr[i] = T(i)
	}

	// Fill the rest with random values
	for i := (size * 2) / 3; i < size; i++ {
		arr[i] = randomFunc()
	}
	return arr
}
