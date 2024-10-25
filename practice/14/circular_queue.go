package main

import "fmt"

type CircularQueue struct {
	size int
	nodes []interface{}
	head int
	last int
}

func NewQueue(num int) *CircularQueue {
	var circularQueue CircularQueue
	circularQueue = CircularQueue{size: num + 1,head: 0, last: 0}
	circularQueue.nodes = make([]interface{}, circularQueue.size)
	return &circularQueue
}

func (c CircularQueue) IsComplete() bool {
	return c.head == (c.last+1)%c.size
}

func (c *CircularQueue) Add(element interface{}) {
	if c.IsComplete() {
		panic("Queue is Complete Utilized")
	}
	c.nodes[c.last] = element
	c.last = (c.last + 1) % c.size
}

func main() {
	var circularQueue *CircularQueue
	circularQueue = NewQueue(5)

	circularQueue.Add(1)
	circularQueue.Add(2)
	circularQueue.Add(3)
	circularQueue.Add(4)
	circularQueue.Add(5)

	fmt.Println(circularQueue.nodes)
}