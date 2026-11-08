package graph

import (
	"container/heap"

	"skiena/pkg/dsu"
	pq "skiena/pkg/prioque"
)

type pqItem[T comparable] struct {
	e *Edge[T]
}

func (it *pqItem[T]) Priority() int {
	return it.e.Weight
}

func (g *Graph[T]) PrimMST(start *Node[T]) *Graph[T] {
	if g.Directed {
		return nil
	}

	st := NewGraph[T](false)
	m := g.NodesCnt - 1
	intree := make(map[*Node[T]]bool, g.NodesCnt)
	pq := pq.New(pq.Min)
	heap.Init(pq)
	pqAddEdge := func(u *Node[T]) {
		intree[u] = true
		for _, e := range g.Edges[u.Label] {
			if !intree[e.Dst] {
				heap.Push(pq, &pqItem[T]{e})
			}
		}
	}

	pqAddEdge(start)

	for pq.Len() > 0 && st.EdgesCnt != m {
		e := heap.Pop(pq).(*pqItem[T]).e
		if intree[e.Dst] {
			continue
		}
		st.InsertEdge(e.Src, e.Dst, e.Weight)
		pqAddEdge(e.Dst)
	}

	return st
}

func (g *Graph[T]) Kruskal() *Graph[T] {
	if g.Directed {
		return nil
	}

	st := NewGraph[T](false)
	m := g.NodesCnt - 1
	dsu := dsu.New(g.NodesAll())
	pq := pq.New(pq.Min)
	for _, dstmap := range g.Edges {
		for _, e := range dstmap {
			pq.Push(&pqItem[T]{e})
		}
	}
	heap.Init(pq)

	for pq.Len() > 0 && st.EdgesCnt != m {
		e := heap.Pop(pq).(*pqItem[T]).e
		u, v := e.Src, e.Dst
		if dsu.Union(u, v) {
			st.InsertEdge(u, v, e.Weight)
		}
	}

	return st
}
