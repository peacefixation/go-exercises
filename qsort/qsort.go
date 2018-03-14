package main

import "fmt"
import "math/rand"

const ARRAY_SIZE int = 50
const MAX_NUM int = 50

func main() {
	// populate an array with random numbers
	xs := make([]int, ARRAY_SIZE)
	for i := 0; i < ARRAY_SIZE; i++ {
		xs[i] = rand.Int() % MAX_NUM
	}

	fmt.Println(xs)
	xs = qsort(xs)
	fmt.Println(xs)
}

func qsort(xs []int) []int {
	if len(xs) < 2 {
		return xs
	}

	left, right := 0, len(xs) - 1
	pivot := left + (right - left) / 2

	// move the pivot value to right
	xs[pivot], xs[right] = xs[right], xs[pivot]

	// move all the elements less than the pivot value to the left
	for i := range xs {
		if xs[i] < xs[right] {
			xs[i], xs[left] = xs[left], xs[i]
			left++
		}
	}

	// move the pivot value to after the last smaller element on the left
	xs[left], xs[right] = xs[right], xs[left]

	// sort the partitions
	qsort(xs[:left])
	qsort(xs[left + 1:])

	return xs
}