package main

import (
	"fmt"

	"skiena/pkg/sort"
)

func main() {
	s1 := []int{1, 3, 4, 6}
	s2 := []int{4, 5, 8, 12}
	z := 14
	found := false
	i, j := 0, len(s2)-1
	x1, x2 := -1, -1

	sort.Quicksort(s1, 0, len(s1)-1)
	sort.Quicksort(s2, 0, len(s2)-1)

	for i < len(s1) && j >= 0 {
		if s1[i]+s2[j] > z {
			j--
		} else if s1[i]+s2[j] < z {
			i++
		} else {
			found = true
			x1, x2 = s1[i], s2[j]
			break
		}
	}

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	if !found {
		fmt.Printf("pair with sum = %d does not exist\n", z)
	} else {
		fmt.Printf("pair with sum = %d is (%d, %d)\n", z, x1, x2)
	}
}
