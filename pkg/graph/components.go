package graph

func (g *Graph[T]) ConnectedComponents() [][]*Node[T] {
	var (
		res       [][]*Node[T]
		component []*Node[T]
	)

	bfs := newBFS(g)
	bfs.onNodeDiscovered = func(n *Node[T]) bool {
		component = append(component, n)
		return true
	}

	for _, n := range g.Nodes {
		if _, ok := bfs.state[n]; ok {
			continue
		}

		bfs.bfs(n)
		res = append(res, component)
		component = []*Node[T]{}
	}

	return res
}

func (g *Graph[T]) FindCutNodes() []*Node[T] {
	if g.Directed {
		return nil
	}

	var res []*Node[T]

	up := make(map[*Node[T]]int)       // earliest accessable time
	tin := make(map[*Node[T]]int)      // entry node time
	children := make(map[*Node[T]]int) // tree edge count
	t := 0

	dfs := newDFS(g)
	dfs.onNodeDiscovered = func(u *Node[T]) bool {
		t++
		up[u] = t
		tin[u] = t

		return true
	}
	dfs.onEdgeBefore = func(e *Edge[T]) bool {
		u, v := e.Src, e.Dst
		if sv, ok := dfs.state[v]; ok && sv == discovered {
			up[u] = min(up[u], tin[v]) // back edge
		}

		return true
	}
	dfs.onEdgeAfter = func(e *Edge[T]) bool {
		u, v := e.Src, e.Dst
		up[u] = min(up[u], up[v]) // tree edge
		if dfs.parent[u] != nil && up[v] >= tin[u] {
			res = append(res, u)
		}
		children[u]++

		return true
	}
	dfs.onNodeProcessed = func(u *Node[T]) bool {
		if dfs.parent[u] == nil && children[u] > 1 {
			res = append(res, u)
		}

		return true
	}

	dfs.runAll()

	return res
}

func (g *Graph[T]) StrongComponents() [][]*Node[T] {
	if !g.Directed {
		return nil
	}

	var (
		res       [][]*Node[T]
		component []*Node[T]
	)

	dfs1Stack := make([]*Node[T], 0, len(g.Nodes))
	dfs1 := newDFS(g)
	dfs1.onNodeProcessed = func(n *Node[T]) bool {
		dfs1Stack = append(dfs1Stack, n)
		return true
	}

	dfs1.runAll()

	gt := g.Transpose()
	dfs2 := newDFS(gt)
	dfs2.onNodeDiscovered = func(n *Node[T]) bool {
		component = append(component, n)
		return true
	}

	for i := len(dfs1Stack) - 1; i >= 0; i-- {
		u := gt.Nodes[dfs1Stack[i].Label] // gt.nodes != g.nodes so dfs2 on gt will fail with g's node

		if _, ok := dfs2.state[u]; ok {
			continue
		}

		dfs2.run(u)
		res = append(res, component)
		component = []*Node[T]{}
	}

	return res
}
