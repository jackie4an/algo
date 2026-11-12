package main

import (
	"fmt"

	"skiena/pkg/graph"
)

func main() {
	g := newGraph()
	parent := make(map[int]int, g.NodeCnt)
	visited := make(map[int]bool, g.NodeCnt)
	path := []int{}

	for n := range g.Nodes {
		if visited[n] {
			continue
		}
		path = bfs(g, n, visited, parent)
		if len(path) > 0 {
			break
		}
	}

	fmt.Println(path)
}

func newGraph() *graph.Graph[int] {
	g := graph.NewGraph[int](false)

	g.InsertEdge(1, 2, 0)
	g.InsertEdge(2, 3, 0)
	g.InsertEdge(3, 4, 0)
	g.InsertEdge(4, 1, 0)

	g.InsertEdge(4, 5, 0)
	g.InsertEdge(5, 6, 0)
	g.InsertEdge(6, 7, 0)
	g.InsertEdge(7, 4, 0)
	g.InsertEdge(1, 7, 0)

	return g
}

func bfs(g *graph.Graph[int], s int, visited map[int]bool, parent map[int]int) []int {
	que := []int{s}
	parent[s] = -1
	visited[s] = true

	for head := 0; head < len(que); head++ {
		u := que[head]
		for v := range g.Edges[u] {
			if visited[v] {
				if parent[u] == parent[v] {
					return []int{u, v, parent[u]}
				} else {
					continue
				}
			}
			parent[v] = u
			visited[v] = true
			que = append(que, v)
		}
	}

	return []int{}
}
