package main

import (
	"fmt"

	"skiena/pkg/graph"
)

func main() {
	// tree
	g := newTree()
	deg := make(map[int]int, g.NodeCnt)
	visited := make(map[int]bool, g.NodeCnt)
	leaves := []int{}
	res := []int{}

	for n, node := range g.Nodes {
		deg[n] = len(node.Neighbors)
		if deg[n] == 1 {
			leaves = append(leaves, n)
		}
	}

	for head := 0; head < len(leaves); head++ {
		leaf := leaves[head]

		if visited[leaf] {
			continue
		}

		p := -1
		for _, v := range g.Nodes[leaf].Neighbors {
			if !visited[v.Label] {
				p = v.Label
				break
			}
		}

		if p == -1 {
			continue
		}

		visited[p] = true
		res = append(res, p)

		for v := range g.Edges[p] {
			if visited[v] {
				continue
			}
			deg[v]--
			if deg[v] == 1 {
				leaves = append(leaves, v)
			}
		}
	}

	fmt.Println(res)
}

func newTree() *graph.Graph[int] {
	g := graph.NewGraph[int](false)

	g.InsertEdge(1, 2, 0)
	g.InsertEdge(1, 3, 0)
	g.InsertEdge(1, 9, 0)
	g.InsertEdge(3, 4, 0)
	g.InsertEdge(3, 5, 0)
	g.InsertEdge(2, 6, 0)
	g.InsertEdge(2, 7, 0)
	g.InsertEdge(2, 8, 0)

	return g
}
