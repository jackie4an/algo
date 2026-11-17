package sort

import (
	"cmp"
)

func SelectionSort[T cmp.Ordered](items []T) {
	for i := range items {
		minIndex := i
		for j := i + 1; j < len(items); j++ {
			if items[j] < items[minIndex] {
				minIndex = j
			}
		}
		items[i], items[minIndex] = items[minIndex], items[i]
	}
}
