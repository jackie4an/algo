package graph

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Node[T comparable] struct {
	Label     T
	Neighbors []*Node[T]
	Incoming  []*Edge[T]
}

func (n *Node[T]) AddNeighbor(neighbor *Node[T]) {
	if !slices.Contains(n.Neighbors, neighbor) {
		n.Neighbors = append(n.Neighbors, neighbor)
	}
}

func (n *Node[T]) AddIncoming(e *Edge[T]) {
	if !slices.Contains(n.Incoming, e) {
		n.Incoming = append(n.Incoming, e)
	}
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("(%v)", n.Label)
}

type Edge[T comparable] struct {
	Src    *Node[T]
	Dst    *Node[T]
	Weight int
}

func (e *Edge[T]) String() string {
	return fmt.Sprintf("(%v)---[w:%d]--->(%v)", e.Src.Label, e.Weight, e.Dst.Label)
}

type Graph[T comparable] struct {
	Nodes    map[T]*Node[T]
	Edges    map[T]map[T]*Edge[T]
	Directed bool
	EdgesCnt int
	NodesCnt int
}

func NewGraph[T comparable](directed bool) *Graph[T] {
	return &Graph[T]{
		Nodes:    make(map[T]*Node[T]),
		Edges:    make(map[T]map[T]*Edge[T]),
		Directed: directed,
	}
}

func (g *Graph[T]) InsertEdge(from, to T, weight int) {
	if g.insertEdge(from, to, weight) {
		g.EdgesCnt++
	}

	if !g.Directed {
		g.insertEdge(from, to, weight)
	}
}

func (g *Graph[T]) AddNode(label T) *Node[T] {
	if g.Nodes == nil {
		g.Nodes = make(map[T]*Node[T])
	}

	n, ok := g.Nodes[label]
	if !ok {
		n = &Node[T]{Label: label}
		g.Nodes[label] = n
		g.NodesCnt++
	}

	return n
}

func (g *Graph[T]) Transpose() *Graph[T] {
	if !g.Directed {
		return g
	}

	gt := NewGraph[T](true)

	for e := range g.BFSEdgesAll() {
		gt.InsertEdge(e.Dst.Label, e.Src.Label, e.Weight)
	}

	return gt
}

func (g Graph[T]) String() string {
	var sb strings.Builder

	sb.WriteString("Nodes:\n")
	g.nodesToString(&sb)
	sb.WriteString("\nCnt: ")
	sb.WriteString(strconv.Itoa(g.NodesCnt))
	sb.WriteByte('\n')

	sb.WriteString("Edges:\n")
	g.edgesToString(&sb)
	sb.WriteString("\nCnt: ")
	sb.WriteString(strconv.Itoa(g.EdgesCnt))

	return sb.String()
}

func (g *Graph[T]) insertEdge(from, to T, weight int) bool {
	if g.Edges == nil {
		g.Edges = make(map[T]map[T]*Edge[T])
	}

	if g.Edges[from] == nil {
		g.Edges[from] = make(map[T]*Edge[T])
	}

	if e, ok := g.Edges[from][to]; ok {
		e.Weight = weight
		return false
	}

	src, ok := g.Nodes[from]
	if !ok {
		src = g.AddNode(from)
	}

	dst, ok := g.Nodes[to]
	if !ok {
		dst = g.AddNode(to)
	}

	e := &Edge[T]{
		Src:    src,
		Dst:    dst,
		Weight: weight,
	}

	g.Edges[from][to] = e

	src.AddNeighbor(dst)
	dst.AddIncoming(e)

	return true
}

func (g *Graph[T]) nodesToString(sb *strings.Builder) {
	first := true
	for _, n := range g.Nodes {
		if !first {
			sb.WriteByte('\n')
		}
		sb.WriteString(n.String())
		first = false
	}
}

func (g *Graph[T]) edgesToString(sb *strings.Builder) {
	first := true
	for _, dst := range g.Edges {
		for _, e := range dst {
			if !first {
				sb.WriteByte('\n')
			}
			sb.WriteString(e.String())
			first = false
		}
	}
}
