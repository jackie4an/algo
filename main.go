package main

import (
	"fmt"

	"skiena/pkg/graph"
)

func main() {
	g := graph.NewGraph[int](true)

	// for i := 1; i <= 3; i++ {
	// 	g.InsertEdge(i, i+1, i*i)
	// }
	// g.InsertEdge(2, 1, -1)

	g.InsertEdge(1, 2, 0)
	g.InsertEdge(2, 3, 0)
	g.InsertEdge(2, 4, 0)
	g.InsertEdge(2, 5, 0)
	g.InsertEdge(3, 1, 0)
	g.InsertEdge(4, 1, 0)
	g.InsertEdge(4, 6, 0)
	g.InsertEdge(4, 8, 0)
	g.InsertEdge(5, 6, 0)
	g.InsertEdge(6, 7, 0)
	g.InsertEdge(7, 5, 0)
	g.InsertEdge(8, 6, 0)

	fmt.Println(g.StrongComponents())
}
