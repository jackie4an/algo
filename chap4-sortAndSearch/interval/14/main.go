package main

import (
	"cmp"
	"fmt"
	"slices"
)

type interval struct {
	start, end int
}

func main() {
	intervals := []interval{
		{1, 3}, {2, 6}, {8, 10}, {7, 18},
	}

	slices.SortFunc(intervals, func(a, b interval) int {
		c := cmp.Compare(a.start, b.start)
		if c == 0 {
			c = cmp.Compare(a.end, b.end)
		}
		return c
	})

	n := len(intervals)
	res := []interval{intervals[0]}

	for i := 1; i < n; i++ {
		last := &res[len(res)-1]

		if intervals[i].start < last.end {
			last.end = max(last.end, intervals[i].end)
		} else {
			res = append(res, intervals[i])
		}
	}

	fmt.Println(res)
}
