package main

import (
	"fmt"
	"math"
	"slices"

	"skiena/pkg/graph"
)

type solver struct {
	g       *graph.Graph[int]
	memo    [][][]int
	visited map[int]bool
	dist    map[int]int
	parent  map[int]int
	s, t    int
}

func main() {
	n := 8
	s, t := 1, 4
	slv := newSolver(n, s, t)

	paths := slv.solve()

	for _, p := range paths {
		fmt.Println(p)
	}
}

func newSolver(n, s, t int) *solver {
	slv := &solver{
		g:       graph.NewGraph[int](true),
		memo:    make([][][]int, n+1),
		visited: make(map[int]bool, n),
		dist:    make(map[int]int, n),
		parent:  make(map[int]int, n),
		s:       s,
		t:       t,
	}

	slv.g.InsertEdge(1, 2, 0)
	slv.g.InsertEdge(1, 4, 0)
	slv.g.InsertEdge(1, 3, 0)
	slv.g.InsertEdge(2, 4, 0)
	slv.g.InsertEdge(3, 2, 0)
	slv.g.InsertEdge(3, 4, 0)
	slv.g.InsertEdge(3, 5, 0)
	slv.g.InsertEdge(5, 6, 0)
	slv.g.InsertEdge(5, 7, 0)
	slv.g.InsertEdge(6, 4, 0)
	slv.g.InsertEdge(7, 8, 0)

	for i := range n {
		slv.dist[i+1] = math.MaxInt
	}

	return slv
}

func (s *solver) solve() [][]int {
	return s.bfs(s.s)
}

func (s *solver) findPath(v int) []int {
	var path []int
	for curr := v; curr != -1; curr = s.parent[curr] {
		path = append(path, curr)
	}
	slices.Reverse(path)
	return path
}

// shortest paths
func (s *solver) bfs(start int) [][]int {
	var paths [][]int

	s.visited[start] = true
	s.dist[start] = 0
	s.parent[start] = -1

	que := []int{start}

	for head := 0; head < len(que); head++ {
		u := que[head]
		for v := range s.g.Edges[u] {
			if s.visited[v] {
				continue
			}

			currDist := s.dist[u] + 1

			if s.dist[v] > currDist {
				s.dist[v] = currDist
			}

			if v != s.t {
				s.visited[v] = true
			}

			s.parent[v] = u
			que = append(que, v)

			if v == s.t {
				if s.dist[v] == currDist {
					path := s.findPath(v)
					paths = append(paths, path)
				} else {
					return paths
				}
			}
		}
	}

	return paths
}

// all paths
func (s *solver) dfs(u int) [][]int {
	if s.visited[u] {
		return s.memo[u]
	}

	s.visited[u] = true

	for v := range s.g.Edges[u] {
		oldPath := s.dfs(v)
		path := make([][]int, len(oldPath))
		for i := range oldPath {
			if oldPath[i][0] != s.t {
				path[i] = []int{u}
				continue
			}
			path[i] = append(oldPath[i], u)
		}
		s.memo[u] = append(s.memo[u], path...)
	}

	if len(s.memo[u]) == 0 {
		s.memo[u] = [][]int{{u}}
	}

	return s.memo[u]
}
