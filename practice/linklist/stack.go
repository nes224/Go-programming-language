package main

import (
	"fmt"
	"sync"
)

type MyStack struct {
	items []int
	mu sync.Mutex
}

func (s *MyStack) prepareStack(n int) {
	s.items = append(s.items, n)
}

func (s *MyStack) Pop() {
	s.mu.Lock()
	defer s.mu.Unlock()
	l := len(s.items) - 1
	s.items = s.items[:l]
}

func Stack() {
	originNumber := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myStack := MyStack{}
	fmt.Println(myStack.items)
	for _, v := range originNumber {
		myStack.prepareStack(v)
	}
	fmt.Println(myStack.items)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(){
			wg.Done()
			myStack.Pop()
		}()
	}
	wg.Wait()
	fmt.Println(myStack.items)
}
