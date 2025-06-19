package adjacencylist

import (
	"fmt"
)

type Graph struct {
	AdjacencyList map[int][]int
}

func CreateAdjacencyList() *Graph {
	return &Graph{
		AdjacencyList: make(map[int][]int), //  map[key] value
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.AdjacencyList[u] = append(g.AdjacencyList[u], v)
	g.AdjacencyList[v] = append(g.AdjacencyList[v], u)

}

func (g *Graph) Print() {
	for node, neighbors := range g.AdjacencyList {
		fmt.Printf("%d -> %v\n", node, neighbors)
	}
}

func (g *Graph) RemoveEdge(u, v int) {
	g.AdjacencyList[u] = removeValue(g.AdjacencyList[u], v)
	g.AdjacencyList[v] = removeValue(g.AdjacencyList[v], u) // Omit this for directed graph
}

func removeValue(slice []int, value int) []int {
	newSlice := []int{}
	for _, v := range slice {
		if v != value {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func (g *Graph) BFS_With_Map(start int) {
	visited := make(map[int]bool)
	queue := []int{start}

	fmt.Print("BFS -> ")
	for len(queue) > 0 {

		//DEqueue
		current := queue[0]
		queue = queue[1:]

		//continue if already visited
		if visited[current] {
			continue
		}

		// mark visited
		visited[current] = true
		fmt.Printf("%d ", current)

		//add neighbors to queue
		for _, neighbor := range g.AdjacencyList[current] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
}

func (g *Graph) DFS_MAP(start int, visited map[int]bool) {
	if visited[start] {
		return
	}

	visited[start] = true
	fmt.Printf("%d ", start)

	for _, neighbor := range g.AdjacencyList[start] {
		g.DFS_MAP(neighbor, visited)
	}
}
