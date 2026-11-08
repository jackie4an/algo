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

func NewNode[T comparable](label T) Node[T] {
	return Node[T]{Label: label}
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

func (g *Graph[T]) InsertEdge(src, dst *Node[T], weight int) {
	if g.insertEdge(src, dst, weight) {
		g.EdgesCnt++
	}

	if !g.Directed {
		g.insertEdge(dst, src, weight)
	}
}

func (g *Graph[T]) AddNode(n *Node[T]) *Node[T] {
	if g.Nodes == nil {
		g.Nodes = make(map[T]*Node[T])
	}

	if _, ok := g.Nodes[n.Label]; !ok {
		g.NodesCnt++
	}

	g.Nodes[n.Label] = n

	return n
}

func (g *Graph[T]) Transpose() *Graph[T] {
	if !g.Directed {
		return g
	}

	gt := NewGraph[T](true)

	for e := range g.BFSEdgesAll() {
		src, dst := NewNode(e.Src.Label), NewNode(e.Dst.Label)
		gt.InsertEdge(&src, &dst, e.Weight)
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

func (g *Graph[T]) insertEdge(src, dst *Node[T], weight int) bool {
	srcl, dstl := src.Label, dst.Label

	g.AddNode(src)
	g.AddNode(dst)

	if g.Edges == nil {
		g.Edges = make(map[T]map[T]*Edge[T])
	}

	if g.Edges[srcl] == nil {
		g.Edges[srcl] = make(map[T]*Edge[T])
	}

	if e, ok := g.Edges[srcl][dstl]; ok {
		e.Weight = weight
		return false
	}

	e := &Edge[T]{
		Src:    src,
		Dst:    dst,
		Weight: weight,
	}

	g.Edges[srcl][dstl] = e

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
	for e := range g.BFSEdgesAll() {
		if !first {
			sb.WriteByte('\n')
		}
		sb.WriteString(e.String())
		first = false
	}
}
