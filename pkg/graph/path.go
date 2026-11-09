package graph

import (
	"math"
	"slices"

	"skiena/pkg/prioque"
)

func (g *Graph[T]) FindPath(start, end *Node[T], parent map[*Node[T]]*Node[T]) []*Node[T] {
	var path []*Node[T]

	for curr := end; curr != nil; curr = parent[curr] {
		path = append(path, curr)
	}
	slices.Reverse(path)

	if len(path) > 0 && path[0] != start {
		return nil
	}

	return path
}

func (g *Graph[T]) ShortestPathDijkstra(start, end *Node[T]) []*Node[T] {
	dist := make(map[*Node[T]]int)
	intree := make(map[*Node[T]]bool)
	parent := make(map[*Node[T]]*Node[T])

	for _, n := range g.Nodes {
		dist[n] = math.MaxInt
	}
	dist[start] = 0

	pq := prioque.New[*Node[T]](prioque.Min)
	pqAddNode := func(n *Node[T]) {
		for _, e := range g.Edges[n.Label] {
			if intree[e.Dst] {
				continue
			}
			if dist[n]+e.Weight < dist[e.Dst] {
				dist[e.Dst] = dist[n] + e.Weight
				parent[e.Dst] = n
				pq.Push(&prioque.Item[*Node[T]]{
					Value:    e.Dst,
					Priority: dist[e.Dst],
				})
			}
		}
	}

	pq.Push(&prioque.Item[*Node[T]]{
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

		pqAddNode(next)

		if next == end {
			break
		}
	}

	return g.FindPath(start, end, parent)
}
