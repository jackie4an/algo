package graph_test

import (
	"testing"

	"skiena/pkg/graph"
)

func TestInsertDirectedEdge(t *testing.T) {
	g := graph.NewGraph[int](true)

	g.InsertEdge(1, 2, 3)
	g.InsertEdge(2, 3, 5)
	g.InsertEdge(3, 1, 4)

	for u := 1; u <= 3; u++ {
		if len(g.Nodes[u].Neighbors) != 1 || len(g.Nodes[u].Incoming) != 1 {
			t.Errorf("expexted neighbors and incoming length = 1, got neighbors: %d, incoming: %d",
				len(g.Nodes[u].Neighbors), len(g.Nodes[u].Incoming))
		}
	}

	g.InsertEdge(3, 1, 8)
	if g.Edges[3][1].Weight != 8 {
		t.Errorf("expexted weight for (3, 1) = 8, got %d", g.Edges[3][1].Weight)
	}

	if g.NodeCnt != 3 {
		t.Errorf("expexted node count = 3, got %d", g.NodeCnt)
	}

	if len(g.Nodes) != 3 {
		t.Errorf("expexted node count = 3, got %d, nodes: %v", len(g.Nodes), g.Nodes)
	}

	if g.EdgeCnt != 3 {
		t.Errorf("expexted edge count = 3, got %d", g.EdgeCnt)
	}

	if len(g.Edges) != 3 {
		t.Errorf("expexted edge count = 3, got %d, edges: %v", len(g.Edges), g.Edges)
	}
}

func TestInsertUndirectedEdge(t *testing.T) {
	g := graph.NewGraph[int](false)

	g.InsertEdge(1, 2, 3)
	g.InsertEdge(2, 3, 5)
	g.InsertEdge(3, 1, 4)

	for u := 1; u <= 3; u++ {
		if len(g.Nodes[u].Neighbors) != 2 || len(g.Nodes[u].Incoming) != 2 {
			t.Errorf("expexted neighbors and incoming length = 1, got neighbors: %d, incoming: %d",
				len(g.Nodes[u].Neighbors), len(g.Nodes[u].Incoming))
		}
	}

	g.InsertEdge(3, 1, 8)
	if g.Edges[3][1].Weight != 8 {
		t.Errorf("expexted weight for (3, 1) = 8, got %d", g.Edges[3][1].Weight)
	}

	if g.NodeCnt != 3 {
		t.Errorf("expexted node count = 3, got %d", g.NodeCnt)
	}

	if len(g.Nodes) != 3 {
		t.Errorf("expexted node count = 3, got %d, nodes: %v", len(g.Nodes), g.Nodes)
	}

	if g.EdgeCnt != 3 {
		t.Errorf("expexted edge count = 3, got %d", g.EdgeCnt)
	}

	allEdges := 0
	for _, vs := range g.Edges {
		for range vs {
			allEdges++
		}
	}
	if allEdges != 6 {
		t.Errorf("expexted all edges count = 6, got %d, edges: %v", allEdges, g.Edges)
	}
}

func TestTranspose(t *testing.T) {
	g := graph.NewGraph[int](true)

	g.InsertEdge(1, 2, 3)
	g.InsertEdge(2, 3, 5)
	g.InsertEdge(3, 1, 4)

	gt := g.Transpose()

	if gt.NodeCnt != g.NodeCnt {
		t.Errorf("expected equal node count, got g: %d, gt: %d", g.NodeCnt, gt.NodeCnt)
	}

	if gt.EdgeCnt != g.EdgeCnt {
		t.Errorf("expected equal edge count, got g: %d, gt: %d", g.EdgeCnt, gt.EdgeCnt)
	}

	for u, edges := range g.Edges {
		for v := range edges {
			if _, ok := gt.Edges[v][u]; !ok {
				t.Errorf("expexted edge (%d, %d) to be in gt", v, u)
			}
		}
	}

	g.Nodes[1] = nil
	if gt.Nodes[1] == nil {
		t.Errorf("expexted node 1 in gt to be not nil, got nil")
	}
}
