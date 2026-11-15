package graph

import (
	"fmt"
	"math"
	"slices"
)

type NetworkGraph[T comparable] struct {
	Nodes map[T]*Node[T]
	Edges map[T]map[T]*Edge[T] // e.Weight == capacity
	Flow  map[T]map[T]int
}

func NewNetworkGraph[T comparable]() *NetworkGraph[T] {
	return &NetworkGraph[T]{
		Nodes: make(map[T]*Node[T]),
		Edges: make(map[T]map[T]*Edge[T]),
		Flow:  make(map[T]map[T]int),
	}
}

func (g *NetworkGraph[T]) InsertEdge(from, to T, capacity int) {
	if g.Edges == nil {
		g.Edges = make(map[T]map[T]*Edge[T])
	}

	if g.Flow == nil {
		g.Flow = make(map[T]map[T]int)
	}

	if _, exists := g.Edges[from]; !exists {
		g.Edges[from] = make(map[T]*Edge[T])
	}

	if _, exists := g.Edges[to]; !exists {
		g.Edges[to] = make(map[T]*Edge[T])
	}

	if _, exists := g.Flow[from]; !exists {
		g.Flow[from] = make(map[T]int)
	}

	if _, exists := g.Flow[to]; !exists {
		g.Flow[to] = make(map[T]int)
	}

	src, dst := g.AddNode(from), g.AddNode(to)

	forwardEdge := &Edge[T]{
		Src:    from,
		Dst:    to,
		Weight: capacity,
	}

	residualEdge := &Edge[T]{
		Src:    to,
		Dst:    from,
		Weight: 0,
	}

	src.addNeighbor(dst)
	dst.addIncoming(forwardEdge)

	dst.addNeighbor(src)
	src.addIncoming(residualEdge)

	g.Edges[from][to] = forwardEdge
	g.Edges[to][from] = residualEdge
	g.Flow[from][to] = 0
	g.Flow[to][from] = 0
}

func (g *NetworkGraph[T]) AddNode(label T) *Node[T] {
	if g.Nodes == nil {
		g.Nodes = make(map[T]*Node[T])
	}

	n, exists := g.Nodes[label]

	if !exists {
		n = &Node[T]{Label: label}
		g.Nodes[label] = n
	}

	return n
}

func (g *NetworkGraph[T]) NodeMap() map[T]*Node[T] {
	return g.Nodes
}

func (g *NetworkGraph[T]) EdgeMap() map[T]map[T]*Edge[T] {
	return g.Edges
}

func (g *NetworkGraph[T]) IsDirected() bool {
	return true
}

func (g *NetworkGraph[T]) remainingCapacity(u, v T) int {
	return g.Edges[u][v].Weight - g.Flow[u][v]
}

func (g *NetworkGraph[T]) MaxFlow(s, t T) int {
	flow := 0

	for {
		path := g.shortestPath(s, t)
		if len(path) == 0 {
			return flow
		}
		fmt.Println(path)

		delta := g.findBottleneck(path)
		g.pushFlow(path, delta)
		flow += delta
	}
}

func (g *NetworkGraph[T]) shortestPath(s, t T) []T {
	que := []T{s}
	visitedMap := make(map[T]struct{})
	parent := make(map[T]*T)

	visitedMap[s] = struct{}{}

	for head := 0; head < len(que); head++ {
		u := que[head]
		if u == t {
			break
		}

		for v := range g.Edges[u] {
			remCap := g.remainingCapacity(u, v)
			if _, visited := visitedMap[v]; visited || remCap <= 0 {
				continue
			}
			visitedMap[v] = struct{}{}
			parent[v] = &u
			que = append(que, v)
		}
	}

	if _, hasParent := parent[t]; !hasParent {
		return nil
	}

	var path []T
	for curr := &t; curr != nil; curr = parent[*curr] {
		path = append(path, *curr)
	}
	slices.Reverse(path)

	return path
}

func (g *NetworkGraph[T]) findBottleneck(p []T) int {
	minCap := math.MaxInt

	for i := 1; i < len(p); i++ {
		u := p[i-1]
		v := p[i]
		minCap = min(minCap, g.remainingCapacity(u, v))
	}

	return minCap
}

func (g *NetworkGraph[T]) pushFlow(p []T, f int) {
	for i := 1; i < len(p); i++ {
		u := p[i-1]
		v := p[i]
		g.Flow[u][v] += f
		g.Flow[v][u] -= f
	}
}
