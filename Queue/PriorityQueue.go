package queue

import (
	"container/heap"
	"fmt"
)

type PriorityQueueItem struct {
	Value    int // The value of the item
	Priority int // The priority of the item (lower value = higher priority)
	Index    int // The index of the item in the heap (used internally)
}

type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*PriorityQueueItem) // type assertion
	item.Index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // Mark as removed
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Update(item *PriorityQueueItem, value int, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

func CreatePriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{}
	heap.Init(pq)
	return pq
}

// DemoPriorityQueue demonstrates the usage of the PriorityQueue.
func DemoPriorityQueue() {
	fmt.Println("Priority Queue Demo")
	pq := CreatePriorityQueue()

	// Insert elements into the priority queue
	heap.Push(pq, &PriorityQueueItem{Value: 1, Priority: 3})
	heap.Push(pq, &PriorityQueueItem{Value: 2, Priority: 1})
	heap.Push(pq, &PriorityQueueItem{Value: 3, Priority: 2})

	// Print and remove elements in priority order
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*PriorityQueueItem)
		fmt.Printf("Value: %d, Priority: %d\n", item.Value, item.Priority)
	}
}
