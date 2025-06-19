package dijkstra

import (
	"container/heap"
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

// Dijkstra's algorithm implementation in Go
// This implementation finds the shortest paths from a source vertex to all other vertices in a weighted graph.

// Edge represents a weighted edge in the graph
type Edge struct {
	To     int
	Weight int
}

// Graph is an adjacency list representation
type Graph map[int][]Edge

// Dijkstra finds shortest paths from source to all vertices
func Dijkstra_PriorityQueue(g Graph, source int) map[int]int {
	// Initialize distances to all vertices as infinite
	// and distance to source as 0
	dist := make(map[int]int)
	for node := range g {
		dist[node] = math.MaxInt
	}
	dist[source] = 0

	// Priority queue to select the vertex with the smallest distance
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{Node: source, Priority: 0})

	// To keep track of visited nodes
	visited := make(map[int]bool)

	for pq.Len() > 0 {
		// Get the node with the smallest distance
		item := heap.Pop(pq).(*Item)
		u := item.Node

		// If this node has already been visited, skip it
		// This is important to avoid processing the same node multiple times
		if visited[u] {
			continue
		}
		visited[u] = true

		// Relaxation step: update distances to neighbors
		// Iterate through all edges from the current node
		for _, edge := range g[u] {
			v := edge.To
			weight := edge.Weight
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(pq, &Item{Node: v, Priority: dist[v]})
			}
		}
	}
	return dist
}

func Practice(g Graph, source int) map[int]int {
	dist := make(map[int]int)

	for node := range g {
		dist[node] = math.MaxInt
	}
	dist[source] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{Node: source, Priority: 0})

	visited := make(map[int]bool)

	for pq.Len() > 0 {

		item := heap.Pop(pq).(*Item)
		u := item.Node

		if visited[u] {
			continue
		}
		visited[u] = true

		for _, edge := range g[u] {
			v := edge.To
			weight := edge.Weight
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(pq, &Item{Node: v, Priority: dist[v]})
			}
		}
	}
	return dist
}

// Item is something we manage in a priority queue.
type Item struct {
	Node     int
	Priority int
	index    int
}

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

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
