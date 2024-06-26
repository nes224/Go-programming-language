package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	items []int
	mu sync.Mutex
}

func (q *Queue) Enqueue(n int) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, n)
}

func (q *Queue) Dequeue() *Queue {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = q.items[1:]
	return q
}

func main() {
	originNumber := []int{1, 2, 3, 4, 5}
	myQueue := Queue{}
	for i := 0; i < len(originNumber); i++ {
		myQueue.Enqueue(originNumber[i])
	}
	var wg sync.WaitGroup
	fmt.Printf("%v\n", myQueue.items)
	for j := 0 ;j < len(myQueue.items); j++ {
		wg.Add(1)
		go func(n int){
			wg.Done()
			myQueue.Dequeue()
		}(myQueue.items[j])
	}
	wg.Wait()
	fmt.Printf("%v\n", myQueue.items)
}
