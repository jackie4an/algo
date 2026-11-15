package graph

func (g *Graph[T]) ConnectedComponents() [][]T {
	var (
		res       [][]T
		component []T
	)

	bfs := NewBFS(g)
	bfs.OnNodeDiscovered = func(n T) bool {
		component = append(component, n)
		return true
	}

	for n := range g.Nodes {
		if _, ok := bfs.State[n]; ok {
			continue
		}

		bfs.bfs(n)
		res = append(res, component)
		component = []T{}
	}

	return res
}

func (g *Graph[T]) FindCutNodes() []T {
	if g.Directed {
		return nil
	}

	var res []T

	up := make(map[T]int)       // earliest accessable time
	tin := make(map[T]int)      // entry node time
	children := make(map[T]int) // tree edge count
	t := 0

	dfs := NewDFS(g)
	dfs.OnNodeDiscovered = func(u T) bool {
		t++
		up[u] = t
		tin[u] = t

		return true
	}
	dfs.OnEdgeBefore = func(e *Edge[T]) bool {
		u, v := e.Src, e.Dst
		if sv, ok := dfs.State[v]; ok && sv == Discovered {
			up[u] = min(up[u], tin[v]) // back edge
		}

		return true
	}
	dfs.OnEdgeAfter = func(e *Edge[T]) bool {
		u, v := e.Src, e.Dst
		up[u] = min(up[u], up[v]) // tree edge
		if dfs.Parent[&u] != nil && up[v] >= tin[u] {
			res = append(res, u)
		}
		children[u]++

		return true
	}
	dfs.OnNodeProcessed = func(u T) bool {
		if dfs.Parent[&u] == nil && children[u] > 1 {
			res = append(res, u)
		}

		return true
	}

	dfs.RunAll()

	return res
}

func (g *Graph[T]) StrongComponents() [][]T {
	if !g.Directed {
		return nil
	}

	var (
		res       [][]T
		component []T
	)

	dfs1Stack := make([]T, 0, len(g.Nodes))
	dfs1 := NewDFS(g)
	dfs1.OnNodeProcessed = func(n T) bool {
		dfs1Stack = append(dfs1Stack, n)
		return true
	}

	dfs1.RunAll()

	gt := g.Transpose()
	dfs2 := NewDFS(gt)
	dfs2.OnNodeDiscovered = func(n T) bool {
		component = append(component, n)
		return true
	}

	for i := len(dfs1Stack) - 1; i >= 0; i-- {
		u := dfs1Stack[i]

		if _, ok := dfs2.State[u]; ok {
			continue
		}

		dfs2.Run(u)
		res = append(res, component)
		component = []T{}
	}

	return res
}
