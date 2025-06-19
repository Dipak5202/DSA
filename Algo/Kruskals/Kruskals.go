package kruskals

import (
	"fmt"
	"sort"
)

type Edge struct {
	u, v   int
	weight int
}

type Graph struct {
	edges    []Edge
	mst      []Edge
	parent   []int
	vertices int
}

func NewGraph(V int) *Graph {
	// Initialize the parent array for union-find
	// Each vertex is its own parent initially
	parent := make([]int, V)
	for i := 0; i < V; i++ {
		parent[i] = i
	}
	return &Graph{
		edges:    []Edge{},
		mst:      []Edge{},
		parent:   parent,
		vertices: V,
	}
}

func (g *Graph) AddWeightedEdge(u, v, w int) {
	g.edges = append(g.edges, Edge{u, v, w})
}

func (g *Graph) findSet(i int) int {
	/// If the parent of i is itself, then i is the representative of the set
	// Otherwise, recursively find the representative of the set
	if g.parent[i] == i {
		return i
	}
	return g.findSet(g.parent[i])
}

func (g *Graph) unionSet(u, v int) {
	g.parent[u] = v
}

func (g *Graph) Kruskal() {
	sort.Slice(g.edges, func(i, j int) bool {
		return g.edges[i].weight < g.edges[j].weight
	})

	for _, e := range g.edges {
		uRep := g.findSet(e.u)
		vRep := g.findSet(e.v)

		if uRep != vRep {
			g.mst = append(g.mst, e)
			g.unionSet(uRep, vRep)
		}
	}
}

func (g *Graph) PrintMST() {
	fmt.Println("Edge : Weight")
	for _, e := range g.mst {
		fmt.Printf("%d - %d : %d\n", e.u, e.v, e.weight)
	}
}
