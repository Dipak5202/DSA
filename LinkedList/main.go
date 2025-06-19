package main

import (
	dl "DSA/LinkedList/DL"
	ll "DSA/LinkedList/LL"
)

func main() {

	// LinkedListExample()
	DublyLinkedListExample()
}

func DublyLinkedListExample() {
	dl := dl.CreateDoublyLinkedList()

	dl.InsertAtEnd(10)
	dl.InsertAtEnd(20)
	dl.InsertAtEnd(30)

	dl.Traversal()
	dl.TraversalReverse()

	dl.InsertAtStart(40)
	dl.InsertAtStart(50)

	dl.Traversal()
	dl.TraversalReverse()

	dl.InsertBetweenBerfore(111, 30)
	dl.InsertBetweenBerfore(222, 10)
	dl.InsertBetweenBerfore(333, 50)

	dl.Traversal()
	dl.TraversalReverse()

	dl.InsertBetweenAfter(444, 333)
	dl.InsertBetweenAfter(555, 10)
	dl.Traversal()
	dl.TraversalReverse()

}

func LinkedListExample() {

	// Create a linked list
	ll := ll.CreateLinkedList()

	// Insert elements at the end
	ll.InsertAtEnd(10)
	ll.InsertAtEnd(20)
	ll.InsertAtEnd(30)
	ll.InsertAtEnd(40)
	ll.InsertAtEnd(50)
	ll.InsertAtEnd(60)
	ll.InsertAtEnd(70)

	ll.Traversal()

	// Insert elements at the beginning
	ll.InsertAtBeginning(5)

	// Traverse the linked list
	ll.Traversal()

	println(ll.Search(30))
	println(ll.Search(70))

	// Check if the linked list is empty
	if ll.IsEmpty() {
		println("Linked List is empty")
	} else {
		println("Linked List is not empty")
	}

	ll.Delete(70)
	ll.Delete(70)
	ll.Delete(20)

	ll.Traversal()

	ll.AddBeforeNode(999, 50)
	ll.AddBeforeNode(888, 5)
	ll.AddBeforeNode(333, 60)
	ll.AddBeforeNode(222, 100)

	ll.Traversal()

	ll.AddAfterNode(111, 50)
	ll.AddAfterNode(123, 60)
	ll.AddAfterNode(321, 888)

	ll.Traversal()

}
