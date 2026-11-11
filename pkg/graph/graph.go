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

type Edge[T comparable] struct {
	Src    T
	Dst    T
	Weight int
}

type Graph[T comparable] struct {
	Nodes    map[T]*Node[T]
	Edges    map[T]map[T]*Edge[T]
	Directed bool
	EdgeCnt  int
	NodeCnt  int
}

func NewNode[T comparable](label T) *Node[T] {
	return &Node[T]{Label: label}
}

func NewGraph[T comparable](directed bool) *Graph[T] {
	return &Graph[T]{
		Nodes:    make(map[T]*Node[T]),
		Edges:    make(map[T]map[T]*Edge[T]),
		Directed: directed,
	}
}

func NewEdge[T comparable](u, v T, weight int) *Edge[T] {
	return &Edge[T]{
		Src:    u,
		Dst:    v,
		Weight: weight,
	}
}

func (g *Graph[T]) InsertEdge(from, to T, weight int) {
	if g.insertEdge(from, to, weight) {
		g.EdgeCnt++
	}

	if !g.Directed {
		g.insertEdge(to, from, weight)
	}
}

func (g *Graph[T]) AddNode(label T) *Node[T] {
	if g.Nodes == nil {
		g.Nodes = make(map[T]*Node[T])
	}

	n, exists := g.Nodes[label]

	if !exists {
		n = NewNode(label)
		g.Nodes[label] = n
		g.NodeCnt++
	}

	return n
}

func (g *Graph[T]) NodesAll() []*Node[T] {
	nodes := make([]*Node[T], 0, g.NodeCnt)

	for _, n := range g.Nodes {
		nodes = append(nodes, n)
	}

	return nodes
}

func (g *Graph[T]) Labels() []T {
	labels := make([]T, 0, g.NodeCnt)

	for n := range g.Nodes {
		labels = append(labels, n)
	}

	return labels
}

func (g *Graph[T]) Transpose() *Graph[T] {
	if !g.Directed {
		return g
	}

	gt := NewGraph[T](true)

	for _, edges := range g.Edges {
		for _, e := range edges {
			gt.InsertEdge(e.Dst, e.Src, e.Weight)
		}
	}

	return gt
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("(%v)", n.Label)
}

func (e *Edge[T]) String() string {
	return fmt.Sprintf("(%v)---[w:%d]--->(%v)", e.Src, e.Weight, e.Dst)
}

func (g *Graph[T]) String() string {
	var sb strings.Builder

	sb.WriteString("Nodes:\n")
	g.nodesToString(&sb)
	sb.WriteString("\nCnt: ")
	sb.WriteString(strconv.Itoa(g.NodeCnt))
	sb.WriteByte('\n')

	sb.WriteString("Edges:\n")
	g.edgesToString(&sb)
	sb.WriteString("\nCnt: ")
	sb.WriteString(strconv.Itoa(g.EdgeCnt))

	return sb.String()
}

func (n *Node[T]) addNeighbor(neighbor *Node[T]) {
	if !slices.Contains(n.Neighbors, neighbor) {
		n.Neighbors = append(n.Neighbors, neighbor)
	}
}

func (n *Node[T]) addIncoming(e *Edge[T]) {
	if !slices.Contains(n.Incoming, e) {
		n.Incoming = append(n.Incoming, e)
	}
}

func (g *Graph[T]) insertEdge(from, to T, weight int) bool {
	src, dst := g.AddNode(from), g.AddNode(to)

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

	e := &Edge[T]{
		Src:    from,
		Dst:    to,
		Weight: weight,
	}

	g.Edges[from][to] = e

	src.addNeighbor(dst)
	dst.addIncoming(e)

	return true
}

func (g *Graph[T]) nodesToString(sb *strings.Builder) {
	first := true
	for _, n := range g.Nodes {
		if !first {
			sb.WriteByte(' ')
		}
		if n == nil {
			sb.WriteString("nil")
		} else {
			sb.WriteString(n.String())
		}
		first = false
	}
}

func (g *Graph[T]) edgesToString(sb *strings.Builder) {
	first := true
	for e := range g.BFSEdgesAll() {
		if !first {
			sb.WriteByte('\n')
		}
		if e == nil {
			sb.WriteString("nil")
		} else {
			sb.WriteString(e.String())
		}
		first = false
	}
}
