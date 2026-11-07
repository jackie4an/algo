package graph

type nodeState byte

const (
	Discovered nodeState = iota + 1
	Processed
)

type traverse[T comparable] struct {
	g                *Graph[T]
	State            map[*Node[T]]nodeState
	Parent           map[*Node[T]]*Node[T]
	OnNodeDiscovered func(*Node[T]) bool
	OnNodeProcessed  func(*Node[T]) bool
}

type dfs[T comparable] struct {
	traverse[T]
	OnEdgeBefore func(*Edge[T]) bool
	OnEdgeAfter  func(*Edge[T]) bool
}

type bfs[T comparable] struct {
	traverse[T]
	OnProcessEdge func(*Edge[T]) bool
}

func NewDFS[T comparable](g *Graph[T]) dfs[T] {
	d := dfs[T]{
		traverse: traverse[T]{
			g:      g,
			State:  make(map[*Node[T]]nodeState),
			Parent: make(map[*Node[T]]*Node[T]),
		},
	}

	return d
}

func NewBFS[T comparable](g *Graph[T]) bfs[T] {
	b := bfs[T]{
		traverse: traverse[T]{
			g:      g,
			State:  make(map[*Node[T]]nodeState),
			Parent: make(map[*Node[T]]*Node[T]),
		},
	}

	return b
}

func (t *traverse[T]) prepare() {
	if t.State == nil {
		t.State = make(map[*Node[T]]nodeState)
	}

	if t.Parent == nil {
		t.Parent = make(map[*Node[T]]*Node[T])
	}
}

func (d *dfs[T]) Run(start *Node[T]) {
	if d.g == nil || start == nil || d.g.Nodes[start.Label] != start {
		return
	}

	d.prepare()

	d.Parent[start] = nil
	if !d.dfsRec(start) {
		return
	}
}

func (d *dfs[T]) RunAll() {
	if d.g == nil {
		return
	}

	d.prepare()

	for _, n := range d.g.Nodes {
		if _, ok := d.State[n]; ok {
			continue
		}

		d.Parent[n] = nil // component's root
		if !d.dfsRec(n) {
			return
		}
	}
}

func (d *dfs[T]) dfsRec(u *Node[T]) bool {
	pu := d.Parent[u]
	d.State[u] = Discovered

	if d.OnNodeDiscovered != nil && !d.OnNodeDiscovered(u) {
		return false
	}

	for _, v := range u.Neighbors {
		e, hasEdge := d.g.Edges[u.Label][v.Label]

		if !hasEdge {
			continue
		}

		_, hasParent := d.Parent[v]
		if !hasParent {
			d.Parent[v] = u
		}

		sv, isDiscovered := d.State[v]

		if d.g.Directed || (v != pu && sv != Processed) {
			if d.OnEdgeBefore != nil && !d.OnEdgeBefore(e) {
				return false
			}
		}

		if !isDiscovered {
			dfsOK := d.dfsRec(v)

			if d.OnEdgeAfter != nil && !d.OnEdgeAfter(e) {
				return false
			}

			if !dfsOK {
				return false
			}
		}
	}

	d.State[u] = Processed
	if d.OnNodeProcessed != nil && !d.OnNodeProcessed(u) {
		return false
	}

	return true
}

func (b *bfs[T]) Run(start *Node[T]) {
	if b.g == nil || start == nil || b.g.Nodes[start.Label] != start {
		return
	}

	b.prepare()

	b.bfs(start)
}

func (b *bfs[T]) RunAll() {
	if b.g == nil {
		return
	}

	b.prepare()

	for _, n := range b.g.Nodes {
		if _, ok := b.State[n]; ok {
			continue
		}

		b.bfs(n)
	}
}

func (b *bfs[T]) bfs(start *Node[T]) {
	que := []*Node[T]{start}

	b.State[start] = Discovered
	b.Parent[start] = nil

	if b.OnNodeDiscovered != nil && !b.OnNodeDiscovered(start) {
		return
	}

	for head := 0; head < len(que); head++ {
		u := que[head]

		for _, v := range u.Neighbors {
			e, hasEdge := b.g.Edges[u.Label][v.Label]

			if !hasEdge {
				continue
			}

			sv, isDiscovered := b.State[v]

			if !isDiscovered {
				que = append(que, v)
				b.Parent[v] = u
				b.State[v] = Discovered

				if b.OnNodeDiscovered != nil && !b.OnNodeDiscovered(v) {
					return
				}
			}

			if b.g.Directed || sv != Processed {
				if b.OnProcessEdge != nil && !b.OnProcessEdge(e) {
					return
				}
			}
		}

		b.State[u] = Processed
		if b.OnNodeProcessed != nil && !b.OnNodeProcessed(u) {
			return
		}
	}
}
