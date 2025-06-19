package queue

type Queue struct {
	Front    int
	Rear     int
	Elements [100]int // Fixed size for simplicity
}

func (q *Queue) IsEmpty() bool {
	return q.Front == -1
}

func (q *Queue) IsFull() bool {
	return q.Rear == len(q.Elements)-1
}

func CreateQueue() *Queue {
	return &Queue{Front: -1, Rear: -1}
}

func (q *Queue) Enqueue(value int) bool {
	if q.IsFull() {
		return false
	}

	if q.IsEmpty() {
		q.Front = 0
	}
	q.Rear++
	q.Elements[q.Rear] = value
	return true
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return 0
	}

	if q.Front > q.Rear {
		q.Front = -1
		q.Rear = -1
		println("Empty queue")
		return 0
	}
	println("Dequeued value:", q.Elements[q.Front])
	q.Front++
	return q.Elements[q.Front]
}

func (q *Queue) Traversal() bool {
	if q.IsEmpty() {
		return false
	}

	for i := q.Front; i <= q.Rear; i++ {
		println("Queue value:", q.Elements[i])
	}
	return true
}
