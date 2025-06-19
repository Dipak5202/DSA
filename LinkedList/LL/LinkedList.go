package ll

type Node struct {
	value int
	next  *Node
}
type LinkedList struct {
	Head *Node
}

func CreateLinkedList() *LinkedList {
	return &LinkedList{Head: nil}
}
func (ll *LinkedList) IsEmpty() bool {
	return ll.Head == nil
}
func (ll *LinkedList) InsertAtEnd(value int) {
	node := &Node{value: value, next: nil}
	if ll.IsEmpty() {
		ll.Head = node
		return
	}

	currentNode := ll.Head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = node
}

func (ll *LinkedList) InsertAtBeginning(value int) {
	node := &Node{value: value, next: nil}
	if ll.IsEmpty() {
		ll.Head = node
		return
	}

	node.next = ll.Head
	ll.Head = node
}

func (ll *LinkedList) Traversal() {
	if ll.IsEmpty() {
		println("Linked List is empty")
		return
	}

	head := ll.Head
	for head != nil {
		print(" -> ", head.value)
		head = head.next
	}
	println()
}

func (ll *LinkedList) Search(value int) bool {
	if ll.IsEmpty() {
		return false
	}
	head := ll.Head
	for head != nil {
		if head.value == value {
			return true
		}
		head = head.next
	}
	return false
}

func (ll *LinkedList) Delete(value int) {
	if ll.IsEmpty() {
		println("Linked List is empty")
		return
	}

	if ll.Head.value == value {
		ll.Head = ll.Head.next
		return
	}

	head := ll.Head
	var prev *Node
	for head != nil {
		if head.value == value {
			prev.next = head.next
			return
		}
		prev = head
		head = head.next
	}
	println("Value not found in the linked list")
}

func (ll *LinkedList) AddBeforeNode(value int, target int) {
	if ll.IsEmpty() {
		println("Linked List is empty")
		return
	}

	node := &Node{value: value, next: nil}
	if ll.Head.value == target {
		node.next = ll.Head
		ll.Head = node
		return
	}

	head := ll.Head
	var prev *Node
	for head != nil {
		if head.value == target {
			node.next = head
			prev.next = node
			return
		}
		prev = head
		head = head.next
	}
	println("Value not found in the linked list")
}

func (ll *LinkedList) AddAfterNode(value int, target int) {
	if ll.IsEmpty() {
		println("Linked List is empty")
		return
	}

	node := &Node{value: value, next: nil}
	if ll.Head.value == target {
		node.next = ll.Head.next
		ll.Head.next = node
		return
	}

	head := ll.Head
	for head != nil {
		if head.value == target {
			node.next = head.next
			head.next = node
			return
		}
		head = head.next
	}
	println("Value not found in the linked list")
}
