package queue

type CircularQueue struct {
	Front    int
	Rear     int
	Elements [5]int
}

func (q *CircularQueue) IsEmpty() bool {
	return q.Front == -1
}

func (q *CircularQueue) IsFull() bool {
	return (q.Rear+1)%len(q.Elements) == q.Front
}

func CreateCircularQueue() *CircularQueue {
	return &CircularQueue{Front: -1, Rear: -1}
}

func (q *CircularQueue) Enqueue(value int) bool {
	if q.IsFull() {
		return false
	}

	if q.IsEmpty() {
		q.Front = 0
	}
	q.Rear = (q.Rear + 1) % len(q.Elements)
	q.Elements[q.Rear] = value
	return true
}

func (q *CircularQueue) Dequeue() bool {
	if q.IsEmpty() {
		return false
	}

	if q.Front == q.Rear {
		q.Front = -1
		q.Rear = -1
		println("Empty queue")
		return true
	}
	println("Dequeued value:", q.Elements[q.Front])
	q.Front = (q.Front + 1) % len(q.Elements)
	return true
}

func (q *CircularQueue) Traversal() bool {
	if q.IsEmpty() {
		return false
	}

	i := q.Front
	for {
		println("Queue value:", q.Elements[i])
		if i == q.Rear {
			break
		}
		i = (i + 1) % len(q.Elements)
	}
	return true
}
