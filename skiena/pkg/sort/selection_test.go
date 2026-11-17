package sort_test

import (
	"testing"

	"algo/skiena/pkg/sort"
)

func TestSelection(t *testing.T) {
	items := []int{12, 11, 13, 5, 6}
	sort.SelectionSort(items)
}
