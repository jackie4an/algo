package main

import (
	"cmp"
	"fmt"
	"slices"
)

type interval struct {
	start, end int
}

type event struct {
	point, delta int
}

func main() {
	s := []interval{
		{10, 40}, {20, 60}, {11, 90}, {15, 70},
	}
	events := make([]event, 0, 2*len(s))

	for i := range s {
		events = append(events, event{s[i].start, 1})
		events = append(events, event{s[i].end, -1})
	}

	slices.SortFunc(events, func(a, b event) int {
		c := cmp.Compare(a.point, b.point)
		if c == 0 {
			c = max(a.delta, b.delta)
		}
		return c
	})

	curr, best := 0, 0
	bestPoint := -1
	for _, e := range events {
		curr += e.delta
		if curr > best {
			bestPoint = e.point
			best = curr
		}
	}

	fmt.Println(best, bestPoint)
}
