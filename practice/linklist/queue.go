package main

import (
	"fmt"
	"sync"
)

type QueueList struct {
	items []int
	mu    sync.Mutex
}

func (q *QueueList) Enqueue(n int) {
	q.items = append(q.items, n)
}

func (q *QueueList) Dequeue() {
	q.mu.Lock()
	defer q.mu.Unlock()
	l := q.items[1:]
	q.items = l
}

func Queue() {
	originNumber := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myQueue := QueueList{}
	fmt.Println(myQueue.items)
	for _, v := range originNumber {
		myQueue.Enqueue(v)
	}
	fmt.Println(myQueue.items)
	
}
