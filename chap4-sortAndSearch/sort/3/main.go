package main

import (
	"fmt"

	"skiena/pkg/sort"
)

func main() {
	nums := []int{2, 4, 30, 31, 52, 64}
	sorted := sort.MergeSort(nums)
	i, j := 0, len(sorted)-1
	maxSum := sorted[j] + sorted[i]

	for i < j {
		fmt.Printf("(%d, %d)\n", sorted[i], sorted[j])

		if sorted[i]+sorted[j] > maxSum {
			maxSum = sorted[i] + sorted[j]
		}

		i++
		j--
	}

	fmt.Println("max sum:", maxSum)
}
