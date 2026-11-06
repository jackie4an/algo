package graph

type color string

const (
	white color = "male"
	black       = "female"

	uncolored color = "unknown"
)

func (g *Graph[T]) TwoColor() (map[*Node[T]]color, bool) {
	colors := make(map[*Node[T]]color)
	for _, node := range g.Nodes {
		colors[node] = uncolored
	}

	bipartite := true
	bfs := newBFS(g)
	bfs.onProcessEdge = func(e *Edge[T]) bool {
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

	for _, n := range g.Nodes {
		if !bipartite {
			return nil, false
		}

		if _, ok := bfs.state[n]; ok {
			continue
		}

		colors[n] = white
		bfs.bfs(n)
	}

	return colors, true
}
