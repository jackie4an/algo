package graph

import (
	"math"
	"slices"

	"skiena/pkg/prioque"
)

func FindPath[T comparable](start, end T, parent map[*T]*T) []T {
	var path []T

	for curr := &end; curr != nil; curr = parent[curr] {
		path = append(path, *curr)
	}
	slices.Reverse(path)

	if len(path) > 0 && path[0] != start {
		return nil
	}

	return path
}

func (g *Graph[T]) ShortestPathDijkstra(start, end T) ([]T, int) {
	dist, parent := g.dijkstra(start, end)
	return FindPath(start, end, parent), dist[end]
}

func (g *Graph[T]) dijkstra(
	start, end T,
) (map[T]int, map[*T]*T) {
	dist := make(map[T]int)
	intree := make(map[T]bool)
	parent := make(map[*T]*T)

	for n := range g.Nodes {
		dist[n] = math.MaxInt
	}
	dist[start] = 0

	pq := prioque.New[T](prioque.Min)
	inspectOutEdges := func(n T) {
		for _, e := range g.Edges[n] {
			if intree[e.Dst] {
				continue
			}
			if dist[n]+e.Weight < dist[e.Dst] {
				dist[e.Dst] = dist[n] + e.Weight
				parent[&e.Dst] = &n
				pq.Push(&prioque.Item[T]{
					Value:    e.Dst,
					Priority: dist[e.Dst],
				})
			}
		}
	}

	pq.Push(&prioque.Item[T]{
		Value:    start,
		Priority: 0,
	})

	for pq.Len() > 0 {
		nextItem := pq.Pop()
		next := nextItem.Value
		nextDist := nextItem.Priority
		intree[next] = true

		if nextDist > dist[next] {
			continue
		}

		inspectOutEdges(next)

		if next == end {
			break
		}
	}

	return dist, parent
}
