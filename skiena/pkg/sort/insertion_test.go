package sort_test

import (
	"fmt"
	"testing"

	"algo/skiena/pkg/sort"
)

func TestInsertion(t *testing.T) {
	items := []int{12, 11, 13, 5, 6}
	fmt.Println(items)
	sort.InsertionSort(items)
}
