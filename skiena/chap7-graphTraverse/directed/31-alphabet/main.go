package main

import (
	"fmt"
	"slices"
)

type solver struct {
	alphabet map[string]bool
	words    []string
	graph    map[string]map[string]bool
	visited  map[string]bool
	start    string
}

func main() {
	words := []string{
		"QQZ", "QZZ", "XQZ", "XQX", "XXX",
	}
	s := newSolver(words)
	res := s.solve()
	fmt.Println(res)
}

func newSolver(words []string) *solver {
	s := &solver{
		alphabet: make(map[string]bool),
		words:    words,
		graph:    make(map[string]map[string]bool),
		visited:  make(map[string]bool),
		start:    "",
	}

	wordLen := len(words[0])
	for i := 0; i < len(s.words)-1; i++ {
		runes1 := []rune(s.words[i])
		if s.start == "" {
			s.start = string(runes1[0])
		}
		for j := i + 1; j < len(s.words); j++ {
			runes2 := []rune(s.words[j])
			for k := range wordLen {
				str1, str2 := string(runes1[k]), string(runes2[k])
				if runes1[k] == runes2[k] {
					continue
				}
				if _, ok := s.graph[str1]; !ok {
					s.graph[str1] = make(map[string]bool)
				} else if s.graph[str1][str2] {
					break
				}
				s.graph[str1][str2] = true
				s.alphabet[str1] = true
				s.alphabet[str2] = true
				break
			}
		}
	}

	return s
}

func (s *solver) solve() []string {
	var (
		dfs func(u string)
		res []string
	)

	dfs = func(u string) {
		if s.visited[u] {
			return
		}
		s.visited[u] = true
		for v := range s.graph[u] {
			dfs(v)
		}
		res = append(res, u)
	}

	dfs(s.start)
	slices.Reverse(res)

	return res
}
