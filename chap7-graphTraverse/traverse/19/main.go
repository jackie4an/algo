package main

import (
	"fmt"

	"skiena/pkg/graph"
)

type color byte

const (
	white color = iota + 1
	black
	uncolored
)

func main() {
	g := newTree()
	deg := make(map[int]int, g.NodeCnt)
	colors := make(map[int]color, g.NodeCnt)
	parent := make(map[int]int)
	res := []int{}
	leaves := []int{}

	for n, node := range g.Nodes {
		deg[n] = len(node.Neighbors)
		colors[n] = uncolored
		if deg[n] == 1 {
			leaves = append(leaves, n)
			parent[n] = g.Nodes[n].Neighbors[0].Label
		}
	}

	for head := 0; head < len(leaves); head++ {
		leaf := leaves[head]
		p := parent[leaf]

		if colors[leaf] != uncolored {
			deg[p]--
			if deg[p] == 1 {
				for v := range g.Edges[p] {
					if colors[v] != black {
						parent[p] = v
						break
					}
				}
				leaves = append(leaves, p)
			}
			continue
		}

		res = append(res, leaf)
		colors[leaf] = black
		colors[p] = white

		deg[p]--
		if deg[p] == 1 {
			for v := range g.Edges[p] {
				if colors[v] != black {
					parent[p] = v
					break
				}
			}
			leaves = append(leaves, p)
		}
	}

	fmt.Println(res)
	fmt.Println(parent)
}

func newTree() *graph.Graph[int] {
	g := graph.NewGraph[int](false)

	g.InsertEdge(1, 2, 0)
	g.InsertEdge(1, 9, 0)
	g.InsertEdge(1, 0, 0)
	g.InsertEdge(1, 3, 0)
	g.InsertEdge(3, 4, 0)
	g.InsertEdge(3, 5, 0)
	g.InsertEdge(2, 6, 0)
	g.InsertEdge(2, 7, 0)
	g.InsertEdge(2, 8, 0)

	return g
}
