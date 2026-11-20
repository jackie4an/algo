/*
Problem: add edge e = (u, v) with weight w to existing MST T of graph G.
Find MST of G + e with O(n).

Solution: since adding e = (u, v) makes cycle between u and v,
find max(w) in old path from u to v and compare max(w) with e(w)
*/
package main

type solver struct {
	mst map[string]map[string]int
}

func main() {
	s := newSolver()
	s.solve("B", "F", 1)
}

func newSolver() *solver {
	mst := make(map[string]map[string]int, 8)

	mst["A"] = map[string]int{
		"B": 8, "D": 10,
	}
	mst["B"] = map[string]int{
		"A": 8, "C": 9,
	}
	mst["C"] = map[string]int{
		"B": 9, "G": 11,
	}
	mst["D"] = map[string]int{
		"A": 10, "E": 6, "F": 7,
	}
	mst["E"] = map[string]int{
		"D": 6,
	}
	mst["F"] = map[string]int{
		"D": 7,
	}
	mst["G"] = map[string]int{
		"C": 11,
	}

	return &solver{mst}
}

func (s *solver) solve(u, v string, w int) {
	path := s.findPath(u, v)
	maxu, maxv, maxw := s.findMaxEdge(path)

	if maxw <= w {
		return
	}

	delete(s.mst[maxu], maxv)
	delete(s.mst[maxv], maxu)
	s.mst[u][v] = w
	s.mst[v][u] = w
}

func (s *solver) findPath(begin, end string) []string {
	parent := make(map[string]string, len(s.mst))
	visited := make(map[string]bool, len(s.mst))

	var dfs func(u string) bool
	dfs = func(u string) bool {
		visited[u] = true
		for v := range s.mst[u] {
			if !visited[v] {
				parent[v] = u
				if v == end || !dfs(v) {
					return false
				}
			}
		}
		return true
	}

	parent[begin] = ""
	dfs(begin)

	var path []string
	for curr := end; curr != ""; curr = parent[curr] {
		path = append(path, curr)
	}

	return path
}

func (s *solver) findMaxEdge(path []string) (u, v string, w int) {
	w = -1
	for i := 1; i < len(path); i++ {
		currU := path[i-1]
		currV := path[i]
		if s.mst[currU][currV] > w {
			w = s.mst[currU][currV]
			u = currU
			v = currV
		}
	}
	return
}
