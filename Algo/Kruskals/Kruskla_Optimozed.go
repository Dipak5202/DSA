package kruskals

import (
	"fmt"
	"sort"
)

type GraphOptimized struct {
	edges    []Edge
	mst      []Edge
	parent   []int
	rank     []int // Added for union by rank optimization
	vertices int
}

func NewGraphOptimized(V int) *GraphOptimized {
	//Disjoint Set Initialization
	// Using path compression and union by rank optimizations
	parent := make([]int, V)
	rank := make([]int, V)
	for i := 0; i < V; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &GraphOptimized{
		edges:    []Edge{},
		mst:      []Edge{},
		parent:   parent,
		rank:     rank,
		vertices: V,
	}
}

func (g *GraphOptimized) AddWeightedEdge(u, v, w int) {
	g.edges = append(g.edges, Edge{u, v, w})
}

// Path compression optimization
// finding ultimate parent of a node
// This optimization helps flatten the structure of the tree whenever findSet is called
func (g *GraphOptimized) findSet(i int) int {
	if g.parent[i] != i {
		g.parent[i] = g.findSet(g.parent[i])
	}
	return g.parent[i]
}

// Union by rank optimization
func (g *GraphOptimized) unionSet(u, v int) {
	uRoot := g.findSet(u)
	vRoot := g.findSet(v)
	if uRoot == vRoot {
		return
	}
	if g.rank[uRoot] < g.rank[vRoot] {
		g.parent[uRoot] = vRoot
	} else if g.rank[uRoot] > g.rank[vRoot] {
		g.parent[vRoot] = uRoot
	} else {
		g.parent[vRoot] = uRoot
		g.rank[uRoot]++
	}
}

func (g *GraphOptimized) KruskalOp() {
	// Sort edges based on weight
	// This is crucial for Kruskal's algorithm to ensure we always consider the smallest edge
	sort.Slice(g.edges, func(i, j int) bool {
		return g.edges[i].weight < g.edges[j].weight
	})

	for _, e := range g.edges {
		uRep := g.findSet(e.u)
		vRep := g.findSet(e.v)

		// If u and v are in different sets, add the edge to the MST
		// and union the sets
		// This ensures that we only add edges that do not form a cycle
		// in the MST, thus maintaining the properties of a minimum spanning tree.
		if uRep != vRep {
			g.mst = append(g.mst, e)
			g.unionSet(uRep, vRep)
		}
	}
}

func (g *GraphOptimized) PrintMST() {
	fmt.Println("Edge : Weight")
	for _, e := range g.mst {
		fmt.Printf("%d - %d : %d\n", e.u, e.v, e.weight)
	}
}
