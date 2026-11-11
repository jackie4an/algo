package graph

type color string

const (
	white color = "male"
	black       = "female"

	uncolored color = "unknown"
)

func (g *Graph[T]) TwoColor() (map[T]color, bool) {
	colors := make(map[T]color)
	for n := range g.Nodes {
		colors[n] = uncolored
	}

	bipartite := true
	bfs := NewBFS(g)
	bfs.OnProcessEdge = func(e *Edge[T]) bool {
		u, v := e.Src, e.Dst

		if colors[u] == colors[v] {
			bipartite = false
			return false
		}

		switch colors[u] {
		case white:
			colors[v] = black
		case black:
			colors[v] = white
		default:
			colors[v] = uncolored
		}

		return true
	}

	for n := range g.Nodes {
		if !bipartite {
			return nil, false
		}

		if _, ok := bfs.State[n]; ok {
			continue
		}

		colors[n] = white
		bfs.bfs(n)
	}

	return colors, true
}
