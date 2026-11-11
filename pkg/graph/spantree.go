package graph

import (
	"skiena/pkg/dsu"
	"skiena/pkg/prioque"
)

func (g *Graph[T]) PrimMST(start T) *Graph[T] {
	if g.Directed {
		return nil
	}

	st := NewGraph[T](false)
	m := g.NodeCnt - 1
	intree := make(map[T]bool, g.NodeCnt)
	pq := prioque.New[*Edge[T]](prioque.Min)

	pqAddEdge := func(u T) {
		intree[u] = true
		for _, e := range g.Edges[u] {
			if !intree[e.Dst] {
				pq.Push(&prioque.Item[*Edge[T]]{
					Value:    e,
					Priority: e.Weight,
				})
			}
		}
	}

	pqAddEdge(start)

	for pq.Len() > 0 && st.EdgeCnt != m {
		e := pq.Pop().Value
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
	m := g.NodeCnt - 1
	dsu := dsu.New(g.Labels())

	items := make([]*prioque.Item[*Edge[T]], 0, g.EdgeCnt*2)
	for _, dstmap := range g.Edges {
		for _, e := range dstmap {
			items = append(items, &prioque.Item[*Edge[T]]{
				Value:    e,
				Priority: e.Weight,
			})
		}
	}
	pq := prioque.New(prioque.Min, items...)

	for pq.Len() > 0 && st.EdgeCnt != m {
		e := pq.Pop().Value
		u, v := e.Src, e.Dst
		if dsu.Union(u, v) {
			st.InsertEdge(u, v, e.Weight)
		}
	}

	return st
}
