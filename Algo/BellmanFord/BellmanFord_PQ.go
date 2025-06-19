package bellmanford

import (
	"container/heap"
	"math"
)

type Item struct {
	Node     int
	Priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Priority < pq[j].Priority }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

type Edge struct {
	From   int
	To     int
	Weight int
}

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

// BellmanFord_PriorityQueue finds shortest paths from source to all vertices using a priority queue (for educational purposes)
// Returns the distance map and a bool indicating if a negative cycle exists
func BellmanFord_PriorityQueue(n int, edges []Edge, source int) (map[int]int, bool) {
	dist := make(map[int]int)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt
	}
	dist[source] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{Node: source, Priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		u := item.Node
		for _, e := range edges {
			if e.From == u {
				if dist[u] != math.MaxInt && dist[u]+e.Weight < dist[e.To] {
					dist[e.To] = dist[u] + e.Weight
					heap.Push(pq, &Item{Node: e.To, Priority: dist[e.To]})
				}
			}
		}
	}

	// Negative cycle detection: one more relaxation over all edges
	for _, e := range edges {
		if dist[e.From] != math.MaxInt && dist[e.From]+e.Weight < dist[e.To] {
			return dist, true // Negative cycle detected
		}
	}
	return dist, false
}
