package main

import (
	adjacencylist "DSA/Graph/Adjacency_List"
	adjacencymatrix "DSA/Graph/Adjacency_Matrix"
	"fmt"
)

func main() {
	// Call the function to create and display the adjacency matrix
	// AdjacencyMatrix()
	// AdjacencyList()
	// AdjacencyLisUsingLinkedList()
	// AdjacencyList_BFS()
	fmt.Println("BFS MAP")
	AdjacencyList_BFS()
	fmt.Println("\n DFS MAP")
	AdjacencyList_DFS()
	fmt.Println("\n BFS LL")
	AdjacencyLisUsingLinkedList_BFS_LL()
	fmt.Println("\n DFS LL")
	AdjacencyLisUsingLinkedList_DFS_LL()
}

func AdjacencyLisUsingLinkedList() {
	adjlist := adjacencylist.CreateAdjacencyListUsingLinkedList(5)

	adjlist.AddEdge(1, 2)
	adjlist.AddEdge(4, 2)
	adjlist.AddEdge(1, 3)
	adjlist.AddEdge(4, 1)
	adjlist.AddEdge(4, 2)
	adjlist.AddEdge(3, 4)

	adjlist.Print()
	// adjlist.Display()

}

func AdjacencyLisUsingLinkedList_BFS_LL() {
	adjlist := adjacencylist.CreateAdjacencyListUsingLinkedList(8)

	adjlist.AddEdge(1, 2)
	adjlist.AddEdge(1, 3)
	adjlist.AddEdge(2, 4)
	adjlist.AddEdge(3, 4)
	adjlist.AddEdge(3, 7)
	adjlist.AddEdge(4, 5)
	adjlist.AddEdge(4, 6)

	// adjlist.Print()
	// adjlist.Display()
	adjlist.BFS_LL(1)

}

func AdjacencyLisUsingLinkedList_DFS_LL() {
	adjlist := adjacencylist.CreateAdjacencyListUsingLinkedList(8)

	adjlist.AddEdge(1, 2)
	adjlist.AddEdge(1, 3)
	adjlist.AddEdge(2, 4)
	adjlist.AddEdge(3, 4)
	adjlist.AddEdge(3, 7)
	adjlist.AddEdge(4, 5)
	adjlist.AddEdge(4, 6)

	// adjlist.Print()
	// adjlist.Display()

	visited := make(map[int]bool)
	fmt.Print("DFS -> ")
	adjlist.DFS_LL(1, visited)

}

func AdjacencyList() {
	adjlist := adjacencylist.CreateAdjacencyList()

	adjlist.AddEdge(1, 2)
	adjlist.AddEdge(5, 3)
	adjlist.AddEdge(2, 4)
	adjlist.AddEdge(4, 6)

	adjlist.Print()
}

func AdjacencyList_BFS() {
	adjlist := adjacencylist.CreateAdjacencyList()

	adjlist.AddEdge(1, 2)
	adjlist.AddEdge(1, 3)
	adjlist.AddEdge(2, 4)
	adjlist.AddEdge(3, 4)
	adjlist.AddEdge(3, 7)
	adjlist.AddEdge(4, 5)
	adjlist.AddEdge(4, 6)

	// adjlist.Print()

	adjlist.BFS_With_Map(1)

}

func AdjacencyList_DFS() {
	adjlist := adjacencylist.CreateAdjacencyList()

	adjlist.AddEdge(1, 2)
	adjlist.AddEdge(1, 3)
	adjlist.AddEdge(2, 4)
	adjlist.AddEdge(3, 4)
	adjlist.AddEdge(3, 7)
	adjlist.AddEdge(4, 5)
	adjlist.AddEdge(4, 6)

	// adjlist.Print()
	visited := make(map[int]bool)
	fmt.Print("DFS -> ")
	adjlist.DFS_MAP(1, visited)

}
func AdjacencyMatrix() {
	// Create an adjacency matrix for a graph with 5 vertices
	adjacencyMatrix := adjacencymatrix.CreateAdjacencyMatrix(4)

	// Add edges to the graph
	adjacencymatrix.AddEdge(adjacencyMatrix, 0, 1)
	adjacencymatrix.AddEdge(adjacencyMatrix, 0, 2)
	adjacencymatrix.AddEdge(adjacencyMatrix, 0, 3)
	adjacencymatrix.AddEdge(adjacencyMatrix, 1, 0)
	adjacencymatrix.AddEdge(adjacencyMatrix, 1, 2)
	adjacencymatrix.AddEdge(adjacencyMatrix, 2, 0)
	adjacencymatrix.AddEdge(adjacencyMatrix, 2, 1)
	adjacencymatrix.AddEdge(adjacencyMatrix, 3, 0)

	// Display the adjacency matrix
	adjacencymatrix.DisplayAdjacencyMatrix(adjacencyMatrix)

	// Check if there is an edge between vertices 0 and 1
	if adjacencymatrix.IsEdge(adjacencyMatrix, 0, 1) {
		println("There is an edge between 0 and 1")
	} else {
		println("There is no edge between 0 and 1")
	}

	// Get the neighbors of vertex 0
	neighbors := adjacencymatrix.GetNeighbors(adjacencyMatrix, 0)
	print("Neighbors of vertex 0: ")
	for _, neighbor := range neighbors {
		print(neighbor, " ")
	}
	println()

	// Get the degree of vertex 0
	degree := adjacencymatrix.GetDegree(adjacencyMatrix, 0)
	println("Degree of vertex 0:", degree)

	// get vertices count
	verticesCount := len(adjacencyMatrix)
	println("Vertices count:", verticesCount)
}
