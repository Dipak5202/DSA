package main

import (
	queue "DSA/Queue"
	stack "DSA/Stack"
)

func main() {

	// StackExample()
	// QueueExample()
	// CircularQueueExample()
	queue.DemoPriorityQueue()
}

func CircularQueueExample() {
	// Create a circular queue
	circularQueue := queue.CreateCircularQueue()

	// Enqueue elements into the circular queue
	circularQueue.Enqueue(10)
	circularQueue.Enqueue(20)
	circularQueue.Enqueue(30)
	circularQueue.Enqueue(40)
	circularQueue.Enqueue(50)

	// Peek the front element
	circularQueue.Traversal()

	// Dequeue an element from the circular queue
	circularQueue.Dequeue()

	circularQueue.Enqueue(60) // This will not be added as the queue is full

	// Peek the front element
	circularQueue.Traversal()

	// Check if the circular queue is empty
	if circularQueue.IsEmpty() {
		println("Circular Queue is empty")
	} else {
		println("Circular Queue is not empty")
	}

	// Check if the circular queue is full
	if circularQueue.IsFull() {
		println("Circular Queue is full")
	} else {
		println("Circular Queue is not full")
	}

	queue.DemoPriorityQueue()
}
func QueueExample() {
	// Create a queue
	queue := queue.CreateQueue()

	// Enqueue elements into the queue
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	// Peek the front element
	queue.Traversal()

	// Dequeue an element from the queue
	queue.Dequeue()

	// Peek the front element
	queue.Traversal()

	// Check if the queue is empty
	if queue.IsEmpty() {
		println("Queue is empty")
	} else {
		println("Queue is not empty")
	}

	// Check if the queue is full
	if queue.IsFull() {
		println("Queue is full")
	} else {
		println("Queue is not full")
	}
}

func StackExample() {
	// Create a stack
	stack := stack.CreateStack()

	// Push elements onto the stack
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	// Peek the top element
	stack.Peek()

	// Pop an element from the stack
	stack.Pop()

	// Check if the stack is empty
	if stack.IsEmpty() {
		println("Stack is empty")
	} else {
		println("Stack is not empty")
	}

	// Check if the stack is full
	if stack.IsFull() {
		println("Stack is full")
	} else {
		println("Stack is not full")
	}
}
