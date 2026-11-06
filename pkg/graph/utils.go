package graph

type edgeClass byte

const (
	Tree edgeClass = iota + 1
	Back
	Forward
	Cross
)

func (g *Graph[T]) IsDAG() bool {
	if !g.Directed {
		return false
	} else if _, hasCycle := g.hasDirectedCycle(); hasCycle {
		return false
	} else {
		return true
	}
}

func (g *Graph[T]) HasCycle() ([]*Node[T], bool) {
	if g.Directed {
		return g.hasDirectedCycle()
	} else {
		return g.hasUndirectedCycle()
	}
}

func (g *Graph[T]) TopSort() ([]*Node[T], bool) {
	if !g.IsDAG() {
		return nil, false
	}

	var stack []*Node[T]

	dfs := newDFS(g)
	dfs.onNodeProcessed = func(n *Node[T]) bool {
		stack = append(stack, n)
		return true
	}

	dfs.runAll()

	res := make([]*Node[T], len(stack))
	for i := len(stack) - 1; i >= 0; i-- {
		res[len(stack)-i-1] = stack[i]
	}

	return res, true
}

func (g *Graph[T]) hasDirectedCycle() ([]*Node[T], bool) {
	var (
		cycle    []*Node[T]
		hasCycle bool
	)

	dfs := newDFS(g)

	dfs.onEdgeBefore = func(e *Edge[T]) bool {
		if dfs.state[e.Dst] == discovered {
			hasCycle = true
			cycle = append(cycle, e.Src)
			return false
		}
		return true
	}

	dfs.onEdgeAfter = func(e *Edge[T]) bool {
		if hasCycle {
			cycle = append(cycle, e.Src)
			return false
		}
		return true
	}

	dfs.runAll()

	return cycle, hasCycle
}

func (g *Graph[T]) hasUndirectedCycle() ([]*Node[T], bool) {
	panic("unimplemented")
}

// func edgeClassification[T comparable](data edgeClassData[T]) edgeClass {
// 	switch {
// 	case data.dstParent == data.src:
// 		return Tree
//
// 	case data.dstState == discovered:
// 		return Back
//
// 	case data.dstState == processed:
// 		if data.dstEntryTime > data.srcEntryTime {
// 			return Forward
// 		} else if data.dstEntryTime < data.srcEntryTime {
// 			return Cross
// 		}
// 	}
//
// 	fmt.Printf("Might be self loop: (%v, %v)\n", data.src, data.dst)
//
// 	return 0
// }
