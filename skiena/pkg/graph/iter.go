package graph

import "iter"

func (g *Graph[T]) BFSNodes(start T) iter.Seq[T] {
	return func(yield func(T) bool) {
		bfs := NewBFS(g)
		bfs.OnNodeDiscovered = yield
		bfs.Run(start)
	}
}

func (g *Graph[T]) BFSNodesAll() iter.Seq[T] {
	return func(yield func(T) bool) {
		bfs := NewBFS(g)
		bfs.OnNodeDiscovered = yield
		bfs.RunAll()
	}
}

func (g *Graph[T]) BFSEdges(start T) iter.Seq[*Edge[T]] {
	return func(yield func(*Edge[T]) bool) {
		bfs := NewBFS(g)
		bfs.OnProcessEdge = yield
		bfs.Run(start)
	}
}

func (g *Graph[T]) BFSEdgesAll() iter.Seq[*Edge[T]] {
	return func(yield func(*Edge[T]) bool) {
		bfs := NewBFS(g)
		bfs.OnProcessEdge = yield
		bfs.RunAll()
	}
}

func (g *Graph[T]) DFSNodes(start T) iter.Seq[T] {
	return func(yield func(T) bool) {
		dfs := NewDFS(g)
		dfs.OnNodeDiscovered = yield
		dfs.Run(start)
	}
}

func (g *Graph[T]) DFSEdges(start T) iter.Seq[*Edge[T]] {
	return func(yield func(*Edge[T]) bool) {
		dfs := NewDFS(g)
		dfs.OnEdgeBefore = yield
		dfs.Run(start)
	}
}
