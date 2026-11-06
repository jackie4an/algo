package sort

import (
	"cmp"
)

func Quicksort[T cmp.Ordered](items []T) {
	quicksort(items, 0, len(items)-1)
}

func quicksort[T cmp.Ordered](items []T, l, h int) {
	if l < h {
		p := partition(items, l, h)
		quicksort(items, l, p-1)
		quicksort(items, p+1, h)
	}
}

func partition[T cmp.Ordered](items []T, l, h int) int {
	p := h
	firsthigh := l

	for i := l; i < h; i++ {
		if items[i] < items[p] {
			swap(items, i, firsthigh)
			firsthigh++
		}
	}

	swap(items, p, firsthigh)

	return firsthigh
}

func swap[T cmp.Ordered](items []T, i, j int) {
	items[i], items[j] = items[j], items[i]
}
