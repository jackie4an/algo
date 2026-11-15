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

func (g *Graph[T]) HasCycle() ([]T, bool) {
	if g.Directed {
		return g.hasDirectedCycle()
	} else {
		return g.hasUndirectedCycle()
	}
}

func (g *Graph[T]) TopSort() ([]T, bool) {
	if !g.IsDAG() {
		return nil, false
	}

	var stack []T

	dfs := NewDFS(g)
	dfs.OnNodeProcessed = func(n T) bool {
		stack = append(stack, n)
		return true
	}

	dfs.RunAll()

	res := make([]T, len(stack))
	for i := len(stack) - 1; i >= 0; i-- {
		res[len(stack)-i-1] = stack[i]
	}

	return res, true
}

func (g *Graph[T]) hasDirectedCycle() ([]T, bool) {
	var (
		cycle    []T
		hasCycle bool
	)

	dfs := NewDFS(g)

	dfs.OnEdgeBefore = func(e *Edge[T]) bool {
		if dfs.State[e.Dst] == Discovered {
			hasCycle = true
			cycle = append(cycle, e.Src)
			return false
		}
		return true
	}

	dfs.OnEdgeAfter = func(e *Edge[T]) bool {
		if hasCycle {
			cycle = append(cycle, e.Src)
			return false
		}
		return true
	}

	dfs.RunAll()

	return cycle, hasCycle
}

func (g *Graph[T]) hasUndirectedCycle() ([]T, bool) {
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
