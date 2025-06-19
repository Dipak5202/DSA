package adjacencylist

import (
	"fmt"
)

type Node struct {
	vertex int
	next   *Node
}

type AdjacenncyList struct {
	vertices       int
	Adjacennc_List []*Node
}

func CreateAdjacencyListUsingLinkedList(vertices int) *AdjacenncyList {
	return &AdjacenncyList{
		vertices:       vertices,
		Adjacennc_List: make([]*Node, vertices),
	}
}

func (aj *AdjacenncyList) AddEdge(u, v int) {

	newnodeV := &Node{
		vertex: v,
		next:   nil,
	}
	newnodeU := &Node{
		vertex: u,
		next:   nil,
	}

	// U -> V
	if aj.Adjacennc_List[u] == nil {
		aj.Adjacennc_List[u] = newnodeV
	} else {
		newnodeV.next = aj.Adjacennc_List[u]
		aj.Adjacennc_List[u] = newnodeV
	}

	// V -> U
	if aj.Adjacennc_List[v] == nil {
		aj.Adjacennc_List[v] = newnodeU
	} else {
		newnodeU.next = aj.Adjacennc_List[v]
		aj.Adjacennc_List[v] = newnodeU
	}
}

func (g *AdjacenncyList) Print() {
	for i := 0; i < g.vertices; i++ {
		fmt.Printf("%d -> ", i)
		curr := g.Adjacennc_List[i]
		for curr != nil {
			fmt.Printf("%d -> ", curr.vertex)
			curr = curr.next
		}
		fmt.Println("nil")
	}
}

func (g *AdjacenncyList) RemoveEdge(u, v int) {
	g.Adjacennc_List[u] = removeFromList(g.Adjacennc_List[u], v)
	g.Adjacennc_List[v] = removeFromList(g.Adjacennc_List[v], u)

}
func removeFromList(head *Node, target int) *Node {
	dummy := &Node{next: head}
	prev := dummy
	curr := head

	for curr != nil {
		if curr.vertex == target {
			prev.next = curr.next
			break
		}
		prev = curr
		curr = curr.next
	}
	return dummy.next
}

func (aj *AdjacenncyList) BFS_LL(start int) {

	visited := make(map[int]bool)
	queue := []int{start}
	fmt.Print("BFS -> ")

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}

		visited[current] = true
		fmt.Printf("%d ", current)

		neighbor := aj.Adjacennc_List[current]

		for neighbor != nil {
			if !visited[neighbor.vertex] {
				queue = append(queue, neighbor.vertex)
			}
			neighbor = neighbor.next
		}
	}

}

func (aj *AdjacenncyList) DFS_LL(start int, visited map[int]bool) {
	if visited[start] {
		return
	}

	visited[start] = true
	fmt.Printf("%d ", start)

	neighbor := aj.Adjacennc_List[start]

	for neighbor != nil {
		aj.DFS_LL(neighbor.vertex, visited)
		neighbor = neighbor.next
	}
}
