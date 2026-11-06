package graph

import "iter"

func (g *Graph[T]) BFSNodes(start *Node[T]) iter.Seq[*Node[T]] {
	return func(yield func(*Node[T]) bool) {
		bfs := newBFS(g)
		bfs.onNodeDiscovered = yield
		bfs.run(start)
	}
}

func (g *Graph[T]) BFSNodesAll() iter.Seq[*Node[T]] {
	return func(yield func(*Node[T]) bool) {
		bfs := newBFS(g)
		bfs.onNodeDiscovered = yield
		bfs.runAll()
	}
}

func (g *Graph[T]) BFSEdges(start *Node[T]) iter.Seq[*Edge[T]] {
	return func(yield func(*Edge[T]) bool) {
		bfs := newBFS(g)
		bfs.onProcessEdge = yield
		bfs.run(start)
	}
}

func (g *Graph[T]) BFSEdgesAll() iter.Seq[*Edge[T]] {
	return func(yield func(*Edge[T]) bool) {
		bfs := newBFS(g)
		bfs.onProcessEdge = yield
		bfs.runAll()
	}
}

func (g *Graph[T]) DFSNodes(start *Node[T]) iter.Seq[*Node[T]] {
	return func(yield func(*Node[T]) bool) {
		dfs := newDFS(g)
		dfs.onNodeDiscovered = yield
		dfs.run(start)
	}
}

func (g *Graph[T]) DFSEdges(start *Node[T]) iter.Seq[*Edge[T]] {
	return func(yield func(*Edge[T]) bool) {
		dfs := newDFS(g)
		dfs.onEdgeBefore = yield
		dfs.run(start)
	}
}
