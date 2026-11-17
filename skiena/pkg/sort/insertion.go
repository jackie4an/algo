package sort

import (
	"cmp"
)

func InsertionSort[T cmp.Ordered](items []T) {
	for i := 1; i < len(items); i++ {
		k := i
		for j := i - 1; j >= 0; j-- {
			if items[k] < items[j] {
				items[k], items[j] = items[j], items[k]
				k = j
			} else {
				break
			}
		}
	}
}
