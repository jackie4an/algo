package graph

type nodeState byte

const (
	discovered nodeState = iota + 1
	processed
)

type traverse[T comparable] struct {
	g                *Graph[T]
	state            map[*Node[T]]nodeState
	parent           map[*Node[T]]*Node[T]
	onNodeDiscovered func(*Node[T]) bool
	onNodeProcessed  func(*Node[T]) bool
}

type dfs[T comparable] struct {
	traverse[T]
	onEdgeBefore func(*Edge[T]) bool
	onEdgeAfter  func(*Edge[T]) bool
}

type bfs[T comparable] struct {
	traverse[T]
	onProcessEdge func(*Edge[T]) bool
}

func newDFS[T comparable](g *Graph[T]) dfs[T] {
	d := dfs[T]{
		traverse: traverse[T]{
			g:      g,
			state:  make(map[*Node[T]]nodeState),
			parent: make(map[*Node[T]]*Node[T]),
		},
	}

	return d
}

func newBFS[T comparable](g *Graph[T]) bfs[T] {
	b := bfs[T]{
		traverse: traverse[T]{
			g:      g,
			state:  make(map[*Node[T]]nodeState),
			parent: make(map[*Node[T]]*Node[T]),
		},
	}

	return b
}

func (t *traverse[T]) prepare() {
	if t.state == nil {
		t.state = make(map[*Node[T]]nodeState)
	}

	if t.parent == nil {
		t.parent = make(map[*Node[T]]*Node[T])
	}
}

func (d *dfs[T]) run(start *Node[T]) {
	if d.g == nil || start == nil || d.g.Nodes[start.Label] != start {
		return
	}

	d.prepare()

	d.parent[start] = nil
	if !d.dfsRec(start) {
		return
	}
}

func (d *dfs[T]) runAll() {
	if d.g == nil {
		return
	}

	d.prepare()

	for _, n := range d.g.Nodes {
		if _, ok := d.state[n]; ok {
			continue
		}

		d.parent[n] = nil // component's root
		if !d.dfsRec(n) {
			return
		}
	}
}

func (d *dfs[T]) dfsRec(u *Node[T]) bool {
	pu := d.parent[u]
	d.state[u] = discovered

	if d.onNodeDiscovered != nil && !d.onNodeDiscovered(u) {
		return false
	}

	for _, v := range u.Neighbors {
		e, hasEdge := d.g.Edges[u.Label][v.Label]

		if !hasEdge {
			continue
		}

		_, hasParent := d.parent[v]
		if !hasParent {
			d.parent[v] = u
		}

		sv, isDiscovered := d.state[v]

		if d.g.Directed || (v != pu && sv != processed) {
			if d.onEdgeBefore != nil && !d.onEdgeBefore(e) {
				return false
			}
		}

		if !isDiscovered {
			dfsOK := d.dfsRec(v)

			if d.onEdgeAfter != nil && !d.onEdgeAfter(e) {
				return false
			}

			if !dfsOK {
				return false
			}
		}
	}

	d.state[u] = processed
	if d.onNodeProcessed != nil && !d.onNodeProcessed(u) {
		return false
	}

	return true
}

func (b *bfs[T]) run(start *Node[T]) {
	if b.g == nil || start == nil || b.g.Nodes[start.Label] != start {
		return
	}

	b.prepare()

	b.bfs(start)
}

func (b *bfs[T]) runAll() {
	if b.g == nil {
		return
	}

	b.prepare()

	for _, n := range b.g.Nodes {
		if _, ok := b.state[n]; ok {
			continue
		}

		b.bfs(n)
	}
}

func (b *bfs[T]) bfs(start *Node[T]) {
	que := []*Node[T]{start}

	b.state[start] = discovered
	b.parent[start] = nil

	if b.onNodeDiscovered != nil && !b.onNodeDiscovered(start) {
		return
	}

	for head := 0; head < len(que); head++ {
		u := que[head]

		for _, v := range u.Neighbors {
			e, hasEdge := b.g.Edges[u.Label][v.Label]

			if !hasEdge {
				continue
			}

			sv, isDiscovered := b.state[v]

			if !isDiscovered {
				que = append(que, v)
				b.parent[v] = u
				b.state[v] = discovered

				if b.onNodeDiscovered != nil && !b.onNodeDiscovered(v) {
					return
				}
			}

			if b.g.Directed || sv != processed {
				if b.onProcessEdge != nil && !b.onProcessEdge(e) {
					return
				}
			}
		}

		b.state[u] = processed
		if b.onNodeProcessed != nil && !b.onNodeProcessed(u) {
			return
		}
	}
}
