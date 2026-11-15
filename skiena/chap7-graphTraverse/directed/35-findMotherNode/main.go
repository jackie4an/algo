package main

import "fmt"

type solver struct {
	g       map[int]map[int]struct{}
	visited map[int]bool
	nodeCnt int
}

func main() {
	slv := newSolver()
	slv.solve(1)
}

func newSolver() *solver {
	s := &solver{
		g:       make(map[int]map[int]struct{}),
		visited: make(map[int]bool),
		nodeCnt: 6,
	}

	for i := 1; i <= 5; i++ {
		s.g[i] = make(map[int]struct{})
	}

	s.g[1][3] = struct{}{}
	s.g[1][4] = struct{}{}
	s.g[1][5] = struct{}{}
	s.g[2][4] = struct{}{}
	s.g[3][6] = struct{}{}
	s.g[4][6] = struct{}{}
	s.g[5][1] = struct{}{}
	s.g[5][2] = struct{}{}

	return s
}

func (s *solver) solve(u int) {
	s.isMother(u)
	s.findMother()
}

func (s *solver) isMother(u int) {
	visited := make(map[int]bool)
	que := []int{u}
	visited[u] = true

	for head := 0; head < len(que); head++ {
		u := que[head]
		for v := range s.g[u] {
			if visited[v] {
				continue
			}
			visited[v] = true
			que = append(que, v)
		}
	}

	if len(visited) == s.nodeCnt {
		fmt.Println(u, "is mother node")
	} else {
		fmt.Println(u, "is not mother node")
	}
}

func (s *solver) findMother() {
	visited := make(map[int]bool, s.nodeCnt)
	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = true
		for v := range s.g[u] {
			if !visited[v] {
				dfs(v)
			}
		}
	}

	var candidate int
	for u := range s.g {
		if !visited[u] {
			dfs(u)
			candidate = u
		}
	}

	s.isMother(candidate)
}
