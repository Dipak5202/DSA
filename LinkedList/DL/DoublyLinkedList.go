package dl

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type DoublyLinkedList struct {
	Head *Node
}

func (dl *DoublyLinkedList) IsEmpty() bool {
	return dl.Head == nil
}

func CreateDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{Head: nil}
}

func (dl *DoublyLinkedList) InsertAtEnd(value int) {
	node := &Node{value: value, prev: nil, next: nil}

	if dl.IsEmpty() {
		dl.Head = node
		return
	}

	currentNode := dl.Head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = node
	node.prev = currentNode
}

func (dl *DoublyLinkedList) InsertAtStart(value int) {
	node := &Node{value: value, prev: nil, next: nil}

	if dl.IsEmpty() {
		dl.Head = node
		return
	}

	node.next = dl.Head
	dl.Head.prev = node
	dl.Head = node
}

func (dl *DoublyLinkedList) Traversal() {
	if dl.IsEmpty() {
		println("Dl Empty")
		return
	}

	currentNode := dl.Head

	for currentNode != nil {
		print(" -> ", currentNode.value)
		currentNode = currentNode.next
	}

	println()
}

func (dl *DoublyLinkedList) TraversalReverse() {

	if dl.IsEmpty() {
		println("Dl Empty")
		return
	}

	currentNode := dl.Head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	for currentNode != nil {
		print(" -> ", currentNode.value)
		currentNode = currentNode.prev
	}
	println()
}

func (dl *DoublyLinkedList) InsertBetweenBerfore(value int, target int) {

	node := &Node{value: value, prev: nil, next: nil}

	if dl.IsEmpty() {
		return
	}

	if dl.Head.value == target {
		node.next = dl.Head
		dl.Head.prev = node
		dl.Head = node
		return
	}

	currentNode := dl.Head
	for currentNode != nil {
		if currentNode.value == target {
			node.next = currentNode
			node.prev = currentNode.prev
			if node.prev != nil { // Check to avoid nil pointer dereference
				node.prev.next = node
			}
			currentNode.prev = node
			return
		}
		currentNode = currentNode.next
	}
}

func (dl *DoublyLinkedList) InsertBetweenAfter(value int, target int) {
	node := &Node{value: value, prev: nil, next: nil}

	if dl.IsEmpty() {
		return
	}

	if dl.Head.value == target {
		node.prev = dl.Head
		node.next = dl.Head.next
		dl.Head.next = node
		dl.Head.next.prev = node
		return
	}

	currentNode := dl.Head

	for currentNode != nil {
		if currentNode.value == target {
			node.next = currentNode.next
			node.prev = currentNode
			if node.next != nil { // Check to avoid nil pointer dereference
				node.next.prev = node
			}
			currentNode.next = node
			return
		}
		currentNode = currentNode.next
	}

}
