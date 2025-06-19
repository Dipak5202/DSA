package main

import (
	bellmanford "DSA/Algo/BellmanFord"
	dijkstra "DSA/Algo/Dijkastra"
	kruskals "DSA/Algo/Kruskals"
	"fmt"
)

func main() {

	KruskalMSTExample()
}

func KruskalMSTExample() {
	g := kruskals.NewGraph(6)
	g.AddWeightedEdge(0, 1, 4) //
	g.AddWeightedEdge(0, 2, 4) //
	g.AddWeightedEdge(1, 2, 2) //
	g.AddWeightedEdge(1, 0, 4) //
	g.AddWeightedEdge(2, 0, 4) //
	g.AddWeightedEdge(2, 1, 2) //
	g.AddWeightedEdge(2, 3, 3) //
	g.AddWeightedEdge(2, 5, 2) //
	g.AddWeightedEdge(2, 4, 4) //
	g.AddWeightedEdge(3, 2, 3) //
	g.AddWeightedEdge(3, 4, 3) //
	g.AddWeightedEdge(4, 2, 4)
	g.AddWeightedEdge(4, 3, 3) //
	g.AddWeightedEdge(5, 2, 2) //
	g.AddWeightedEdge(5, 4, 3) //

	g.Kruskal()
	g.PrintMST()
}

func BellmanFordExampleAndDijkastra() {
	// Example graph: adjacency list
	// 0 --(4)--> 1
	// 0 --(1)--> 2
	// 2 --(2)--> 1
	// 1 --(1)--> 3
	// 2 --(5)--> 3
	graph := dijkstra.Graph{
		0: {{To: 1, Weight: 4}, {To: 2, Weight: 1}},
		1: {{To: 3, Weight: 1}},
		2: {{To: 1, Weight: 2}, {To: 3, Weight: 5}},
		3: {},
	}

	source := 0

	println("Dijkstra's Algorithm using Priority Queue")
	distances := dijkstra.Practice(graph, source)
	fmt.Printf("Shortest distances from node %d:\n", source)
	for node, dist := range distances {
		fmt.Printf("Node %d: %d\n", node, dist)
	}

	println("Dijkstra's Algorithm using Priority Queue - Practice")
	distancesP := dijkstra.Dijkstra_PriorityQueue(graph, source)
	fmt.Printf("Shortest distances from node %d:\n", source)
	for node, dist := range distancesP {
		fmt.Printf("Node %d: %d\n", node, dist)
	}

	// Example edge list for Bellman-Ford
	n := 5 // number of vertices (0 to 4)
	edges := []bellmanford.Edge{
		{From: 0, To: 1, Weight: 6},
		{From: 0, To: 2, Weight: 7},
		{From: 1, To: 2, Weight: 8},
		{From: 1, To: 3, Weight: 5},
		{From: 1, To: 4, Weight: -4},
		{From: 2, To: 3, Weight: -3},
		{From: 2, To: 4, Weight: 9},
		{From: 3, To: 1, Weight: -2},
		{From: 4, To: 0, Weight: 2},
		{From: 4, To: 3, Weight: 7},
	}
	source = 0

	fmt.Println("Bellman-Ford (without priority queue):")
	dist, negCycle := bellmanford.BellmanFord(n, edges, source)
	if negCycle {
		fmt.Println("Graph contains a negative-weight cycle.")
	} else {
		fmt.Printf("Shortest distances from node %d:\n", source)
		for v := 0; v < n; v++ {
			fmt.Printf("Node %d: %d\n", v, dist[v])
		}
	}

	fmt.Println("\nBellman-Ford (with priority queue):")
	distPQ, negCyclePQ := bellmanford.BellmanFord_PriorityQueue(n, edges, source)
	if negCyclePQ {
		fmt.Println("Graph contains a negative-weight cycle.")
	} else {
		fmt.Printf("Shortest distances from node %d:\n", source)
		for v := 0; v < n; v++ {
			fmt.Printf("Node %d: %d\n", v, distPQ[v])
		}
	}
}
