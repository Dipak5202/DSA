package bellmanford

import (
	"math"
)

/*
Why Use a Priority Queue in Dijkstra’s Algorithm?

- Dijkstra’s algorithm always selects the unvisited node with the smallest known distance.
- A priority queue (min-heap) efficiently retrieves and updates the node with the smallest distance in O(log n) time.
- This ensures that each step processes the closest node next, which is essential for Dijkstra’s correctness and efficiency.

Why NOT Use a Priority Queue in Bellman-Ford?

- Bellman-Ford works by relaxing all edges repeatedly, regardless of the order.
- The algorithm does not depend on processing nodes in order of their current shortest distance.
- Using a priority queue does not improve Bellman-Ford’s performance or correctness, because every edge must be checked in every iteration anyway.
- Bellman-Ford’s strength is handling negative weights and detecting negative cycles, which a priority queue does not help with.

Summary:
- Dijkstra’s: Needs a priority queue to always process the closest node next.
- Bellman-Ford: Does not benefit from a priority queue, as it must relax all edges in each iteration.
*/

// Edge represents a weighted edge in the graph
// From -> To with Weight
// Useful for Bellman-Ford's edge list

// BellmanFord finds shortest paths from source to all vertices
// Returns a map of distances and a bool indicating if a negative cycle exists
func BellmanFord(n int, edges []Edge, source int) (map[int]int, bool) {
	dist := make(map[int]int)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt
	}
	dist[source] = 0

	// Relax all edges n-1 times
	for i := 0; i < n-1; i++ {
		for _, e := range edges {
			if dist[e.From] != math.MaxInt && dist[e.From]+e.Weight < dist[e.To] {
				dist[e.To] = dist[e.From] + e.Weight
			}
		}
	}

	// Check for negative-weight cycles
	for _, e := range edges {
		if dist[e.From] != math.MaxInt && dist[e.From]+e.Weight < dist[e.To] {
			return dist, true // Negative cycle detected
		}
	}
	return dist, false
}
