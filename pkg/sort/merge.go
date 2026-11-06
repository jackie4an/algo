package sort

import (
	"cmp"
)

func MergeSort[T cmp.Ordered](items []T) []T {
	if len(items) <= 1 {
		return items
	}

	mid := len(items) / 2

	left := MergeSort(items[:mid])
	right := MergeSort(items[mid:])

	items = merge(left, right)

	return items
}

func merge[T cmp.Ordered](left, right []T) []T {
	result := make([]T, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
