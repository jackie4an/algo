package main

import (
	"fmt"
	"math"
	"slices"

	"skiena/pkg/graph"
)

/*
chutes and ladders
12	11	10
7		8		9
6		5		4
1		2		3

dice sides: 1, 2 or 3
min throws: 1-2-6, 6-9, 9-12
*/

func main() {
	diceSides := 3
	n := 12
	g := initGraph(diceSides, n)

	fmt.Println(bfs(g, diceSides, n))
}

func initGraph(k, n int) *graph.Graph[int] {
	g := graph.NewGraph[int](true)

	for u := 1; u < n; u++ {
		n := min(u+k, n)
		for v := u + 1; v <= n; v++ {
			g.InsertEdge(u, v, 0)
		}
	}

	// ladders
	g.InsertEdge(2, 6, 1)
	g.InsertEdge(5, 10, 1)

	// chutes
	g.InsertEdge(7, 5, -1)
	g.InsertEdge(11, 9, -1)

	return g
}

func bfs(g *graph.Graph[int], k, n int) []int {
	que := []int{1}
	visited := make(map[int]bool, n)
	parent := make(map[int]int, n)
	dist := make(map[int]int, n)
	jump := jump(g, n)

	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt
	}

	visited[1] = true
	parent[1] = -1
	dist[1] = 0

	for head := 0; head < len(que); head++ {
		u := que[head]
		m := min(u+k, n)
		for v := u + 1; v <= m; v++ {
			x := jump[v]

			if visited[x] {
				continue
			}

			visited[x] = true

			if dist[u]+1 < dist[x] {
				dist[x] = dist[u] + 1
				parent[x] = u
			}

			que = append(que, x)
		}
	}

	var path []int
	for curr := n; curr != -1; curr = parent[curr] {
		path = append(path, curr)
	}
	slices.Reverse(path)

	return path
}

func jump(g *graph.Graph[int], n int) map[int]int {
	jump := make(map[int]int)

	for u := 1; u <= n; u++ {
		jump[u] = u
		for v, e := range g.Edges[u] {
			if e.Weight != 0 {
				jump[u] = v
				break
			}
		}
	}

	return jump
}
