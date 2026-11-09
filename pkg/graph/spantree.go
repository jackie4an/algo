package graph

import (
	"skiena/pkg/dsu"
	"skiena/pkg/prioque"
)

func (g *Graph[T]) PrimMST(start *Node[T]) *Graph[T] {
	if g.Directed {
		return nil
	}

	st := NewGraph[T](false)
	m := g.NodesCnt - 1
	intree := make(map[*Node[T]]bool, g.NodesCnt)
	pq := prioque.New[*Edge[T]](prioque.Min)

	pqAddEdge := func(u *Node[T]) {
		intree[u] = true
		for _, e := range g.Edges[u.Label] {
			if !intree[e.Dst] {
				pq.Push(&prioque.Item[*Edge[T]]{
					Value:    e,
					Priority: e.Weight,
				})
			}
		}
	}

	pqAddEdge(start)

	for pq.Len() > 0 && st.EdgesCnt != m {
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
	m := g.NodesCnt - 1
	dsu := dsu.New(g.NodesAll())

	items := make([]*prioque.Item[*Edge[T]], 0, g.EdgesCnt*2)
	for _, dstmap := range g.Edges {
		for _, e := range dstmap {
			items = append(items, &prioque.Item[*Edge[T]]{
				Value:    e,
				Priority: e.Weight,
			})
		}
	}
	pq := prioque.New(prioque.Min, items...)

	for pq.Len() > 0 && st.EdgesCnt != m {
		e := pq.Pop().Value
		u, v := e.Src, e.Dst
		if dsu.Union(u, v) {
			st.InsertEdge(u, v, e.Weight)
		}
	}

	return st
}
