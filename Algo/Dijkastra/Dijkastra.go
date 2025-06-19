package dijkstra

import "math"

// Edge represents a weighted edge in the graph (with From, To, Weight)
type Edge_ struct {
	From   int
	To     int
	Weight int
}

// Dijkstra finds shortest paths from source to all vertices (without priority queue, using edge list)
func Dijkstra(n int, edges []Edge_, source int) map[int]int {
	dist := make(map[int]int)
	visited := make(map[int]bool)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt
	}
	dist[source] = 0

	for i := 0; i < n; i++ {
		u := -1
		minDist := math.MaxInt
		for v := 0; v < n; v++ {
			if !visited[v] && dist[v] < minDist {
				minDist = dist[v]
				u = v
			}
		}
		if u == -1 {
			break
		}
		visited[u] = true
		for _, e := range edges {
			if e.From == u {
				if dist[u] != math.MaxInt && dist[u]+e.Weight < dist[e.To] {
					dist[e.To] = dist[u] + e.Weight
				}
			}
		}
	}
	return dist
}
