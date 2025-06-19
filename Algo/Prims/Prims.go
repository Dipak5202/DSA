package prims

import (
	"math"
)

type Edge struct {
	To     int
	Weight int
}

type Graph map[int][]Edge

// PrimsMST returns the total weight and the edges of the Minimum Spanning Tree
func PrimsMST(g Graph, n int) (int, []struct{ From, To, Weight int }) {
	inMST := make([]bool, n)
	key := make([]int, n)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		key[i] = math.MaxInt
		parent[i] = -1
	}
	key[0] = 0
	totalWeight := 0
	mstEdges := []struct{ From, To, Weight int }{}

	for count := 0; count < n; count++ {
		u := -1
		minKey := math.MaxInt
		for v := 0; v < n; v++ {
			if !inMST[v] && key[v] < minKey {
				minKey = key[v]
				u = v
			}
		}
		if u == -1 {
			break
		}
		inMST[u] = true
		if parent[u] != -1 {
			mstEdges = append(mstEdges, struct{ From, To, Weight int }{parent[u], u, key[u]})
			totalWeight += key[u]
		}
		for _, edge := range g[u] {
			v := edge.To
			w := edge.Weight
			if !inMST[v] && w < key[v] {
				key[v] = w
				parent[v] = u
			}
		}
	}
	return totalWeight, mstEdges
}
