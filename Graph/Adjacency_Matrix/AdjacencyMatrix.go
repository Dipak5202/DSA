// Package adjacencymatrix provides functions to create and manipulate a graph using an adjacency matrix representation.
package adjacencymatrix

// CreateAdjacencyMatrix creates a 2D slice (matrix) to represent the adjacency matrix
// for a graph with the specified number of vertices.
// Each cell in the matrix represents whether an edge exists between two vertices.
// Parameters:
//   - vertices: The number of vertices in the graph.
// Returns:
//   - A 2D slice of booleans representing the adjacency matrix.
func CreateAdjacencyMatrix(vertices int) [][]bool {
	// Create a 2D slice (matrix) to represent the adjacency matrix
	adjacencyMatrix := make([][]bool, vertices)
	for i := range adjacencyMatrix {
		adjacencyMatrix[i] = make([]bool, vertices)
	}
	return adjacencyMatrix
}

// DisplayAdjacencyMatrix prints the adjacency matrix to the console.
// Parameters:
//   - adjacencyMatrix: The 2D slice representing the adjacency matrix.
func DisplayAdjacencyMatrix(adjacencyMatrix [][]bool) {
	// Display the adjacency matrix
	for i := range adjacencyMatrix {
		for j := range adjacencyMatrix[i] {
			if adjacencyMatrix[i][j] {
				print("1 ")
			} else {
				print("0 ")
			}
		}
		println()
	}
}

// AddEdge adds an edge between two vertices in an undirected graph.
// Parameters:
//   - adjacencyMatrix: The 2D slice representing the adjacency matrix.
//   - u: The first vertex.
//   - v: The second vertex.
func AddEdge(adjacencyMatrix [][]bool, u, v int) {
	// Add an edge between vertices u and v (undirected graph)
	adjacencyMatrix[u][v] = true
	adjacencyMatrix[v][u] = true
}

// RemoveEdge removes an edge between two vertices in an undirected graph.
// Parameters:
//   - adjacencyMatrix: The 2D slice representing the adjacency matrix.
//   - u: The first vertex.
//   - v: The second vertex.
func RemoveEdge(adjacencyMatrix [][]bool, u, v int) {
	// Remove an edge between vertices u and v (undirected graph)
	adjacencyMatrix[u][v] = false
	adjacencyMatrix[v][u] = false
}

// IsEdge checks if there is an edge between two vertices in the graph.
// Parameters:
//   - adjacencyMatrix: The 2D slice representing the adjacency matrix.
//   - u: The first vertex.
//   - v: The second vertex.
// Returns:
//   - A boolean indicating whether an edge exists between the two vertices.
func IsEdge(adjacencyMatrix [][]bool, u, v int) bool {
	// Check if there is an edge between vertices u and v
	return adjacencyMatrix[u][v]
}

// GetNeighbors retrieves the neighbors of a given vertex in the graph.
// Parameters:
//   - adjacencyMatrix: The 2D slice representing the adjacency matrix.
//   - vertex: The vertex whose neighbors are to be retrieved.
// Returns:
//   - A slice of integers representing the neighbors of the vertex.
func GetNeighbors(adjacencyMatrix [][]bool, vertex int) []int {
	// Get the neighbors of a vertex
	neighbors := []int{}
	for i, connected := range adjacencyMatrix[vertex] {
		if connected {
			neighbors = append(neighbors, i)
		}
	}
	return neighbors
}

// GetDegree calculates the degree of a given vertex in the graph.
// Parameters:
//   - adjacencyMatrix: The 2D slice representing the adjacency matrix.
//   - vertex: The vertex whose degree is to be calculated.
// Returns:
//   - An integer representing the degree of the vertex.
func GetDegree(adjacencyMatrix [][]bool, vertex int) int {
	// Get the degree of a vertex
	degree := 0
	for _, connected := range adjacencyMatrix[vertex] {
		if connected {
			degree++
		}
	}
	return degree
}

// GetVertices returns the number of vertices in the graph.
// Parameters:
//   - adjacencyMatrix: The 2D slice representing the adjacency matrix.
// Returns:
//   - An integer representing the number of vertices in the graph.
func GetVertices(adjacencyMatrix [][]bool) int {
	// Get the number of vertices in the graph
	return len(adjacencyMatrix)
}
