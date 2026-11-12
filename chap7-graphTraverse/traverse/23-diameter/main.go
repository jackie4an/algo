package main

import (
	"fmt"

	"skiena/pkg/graph"
)

var g *graph.Graph[int]

func init() {
	g = graph.NewGraph[int](false)
	g.InsertEdge(1, 2, 0)
	g.InsertEdge(1, 9, 0)
	g.InsertEdge(1, 0, 0)
	g.InsertEdge(1, 3, 0)
	g.InsertEdge(3, 4, 0)
	g.InsertEdge(3, 5, 0)
	g.InsertEdge(2, 6, 0)
	g.InsertEdge(2, 7, 0)
	g.InsertEdge(2, 8, 0)
}

func main() {
	u, _ := bfs(1)
	v, d := bfs(u)
	fmt.Printf("maxDist(%d, %d) = %d\n", u, v, d)
}

func bfs(s int) (int, int) {
	maxv, dv := s, 0
	dist := make(map[int]int)
	visited := make(map[int]bool)

	dist[s] = 0
	visited[s] = true
	que := []int{s}

	for head := 0; head < len(que); head++ {
		u := que[head]

		for v := range g.Edges[u] {
			if visited[v] {
				continue
			}

			visited[v] = true
			dist[v] = dist[u] + 1
			if dist[v] > dv {
				maxv = v
				dv = dist[v]
			}
			que = append(que, v)
		}
	}

	return maxv, dv
}
