package graph

import (
	"fmt"
)

func (g *Graph[T]) FindPath(start, end *Node[T], parents map[*Node[T]]*Node[T]) {
	if start == end || end == nil {
		fmt.Print(start.Label)
	} else {
		g.FindPath(start, parents[end], parents)
		fmt.Print(" ", end.Label)
	}
}

// func (g *Graph[T]) ShortestPathDijkstra(start, end *Node[T]) []*Node[T] {
// 	dist := make(map[*Node[T]]int)
// 	intree := make(map[*Node[T]]bool)
// 	path := []*Node[T]{start}
//
// 	return path
// }
