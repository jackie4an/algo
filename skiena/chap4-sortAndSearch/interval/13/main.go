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
	timestamp int
	delta     int
}

func main() {
	intervals := []interval{
		{5, 7}, {3, 8}, {2, 3}, {7, 11}, {6, 11},
	}
	events := make([]event, 0, 2*len(intervals))

	for _, ivl := range intervals {
		events = append(events, event{ivl.start, 1})
		events = append(events, event{ivl.end, -1})
	}

	slices.SortFunc(events, func(a, b event) int {
		c := cmp.Compare(a.timestamp, b.timestamp)
		if c == 0 {
			c = cmp.Compare(a.delta, b.delta)
		}
		return c
	})

	curr, best := 0, 0
	bestTime := 0

	for _, e := range events {
		curr += e.delta
		if curr > best {
			best = curr
			bestTime = e.timestamp
		}
	}

	fmt.Println(best, bestTime)
}
