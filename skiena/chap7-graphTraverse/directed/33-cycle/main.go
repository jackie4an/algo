package main

import "fmt"

type solver struct {
	graph    [][]int
	visisted map[int]bool
}

func main() {
	slv := newSolver()
	fmt.Println(slv.solve(1, 5))
}

func newSolver() *solver {
	s := &solver{
		graph:    make([][]int, 9),
		visisted: make(map[int]bool),
	}

	s.graph[1] = []int{2, 5, 8}
	s.graph[2] = []int{1, 3}
	s.graph[3] = []int{2, 4}
	s.graph[4] = []int{3, 5}
	s.graph[5] = []int{1, 4, 6, 7, 8}
	s.graph[6] = []int{5, 7}
	s.graph[7] = []int{5, 6}
	s.graph[8] = []int{1, 5}

	return s
}

func (s *solver) solve(from, to int) []int {
	var (
		dfs   func(u, p int) bool
		cycle []int
	)

	dfs = func(u, p int) bool {
		s.visisted[u] = true
		if u == to {
			cycle = append(cycle, u)
			return false
		}
		for _, v := range s.graph[u] {
			if u == from && v == to ||
				v == from && u == to ||
				v == p ||
				s.visisted[v] {
				continue
			}
			if !dfs(v, u) {
				cycle = append(cycle, u)
				return false
			}
		}
		return true
	}

	dfs(from, -1)

	return cycle
}
