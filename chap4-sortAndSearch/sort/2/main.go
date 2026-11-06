package main

import (
	"fmt"

	"skiena/pkg/sort"
)

func main() {
	unsorted := []int{6, 13, 19, 3, 8}
	sorted := []int{3, 6, 8, 13, 19}

	a(unsorted)
	b(sorted)
	c(unsorted)
	d(sorted)
}

func a(unsorted []int) {
	x := unsorted[0]
	y := unsorted[0]

	for i := 1; i < len(unsorted); i++ {
		if unsorted[i] > x {
			x = unsorted[i]
		}
		if unsorted[i] < y {
			y = unsorted[i]
		}
	}

	fmt.Printf("a). array: %v\n", unsorted)
	fmt.Printf("    x = %d, y = %d\n", x, y)
}

func b(sorted []int) {
	fmt.Printf("b). array: %v\n", sorted)
	fmt.Printf("    x = %d, y = %d\n", sorted[len(sorted)-1], sorted[0])
}

func c(unsorted []int) {
	sorted := sort.MergeSort(unsorted)
	x, y := sorted[len(sorted)-1], sorted[0]
	minDiff := x - y

	for i := 1; i < len(sorted); i++ {
		if sorted[i] == sorted[i-1] {
			continue
		}

		currDiff := sorted[i] - sorted[i-1]
		if currDiff < minDiff {
			minDiff = currDiff
			x, y = sorted[i], sorted[i-1]
		}
	}

	fmt.Printf("c). array: %v\n", unsorted)
	fmt.Printf("    x = %d, y = %d\n", x, y)
}

func d(sorted []int) {
	x, y := sorted[len(sorted)-1], sorted[0]
	minDiff := x - y

	for i := 1; i < len(sorted); i++ {
		if sorted[i] == sorted[i-1] {
			continue
		}

		currDiff := sorted[i] - sorted[i-1]
		if currDiff < minDiff {
			minDiff = currDiff
			x, y = sorted[i], sorted[i-1]
		}
	}

	fmt.Printf("d). array: %v\n", sorted)
	fmt.Printf("    x = %d, y = %d\n", x, y)
}
