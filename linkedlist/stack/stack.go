package main

import (
	"fmt"
	"sync"
)

type Stack struct {
	items []int
	mu sync.Mutex
}

func (s *Stack) Stack(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items,i)
}

func (s *Stack) Pop() []int {
	l := len(s.items) - 1
	s.items = s.items[:l]
	return s.items
}

func main() {
	myStack := Stack{}
	numberOrigin := []int{100,200,300}
	var wg sync.WaitGroup
	fmt.Printf("%v\n",myStack.items)
	for i := 0; i< len(numberOrigin); i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			myStack.Stack(n)
		}(numberOrigin[i])
	}
	wg.Wait()
	fmt.Printf("%v\n",myStack.items)
	for j := 0; j< len(myStack.items); j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			myStack.Pop()
		}()
	}
	wg.Wait()
	fmt.Printf("%v\n",myStack.items)
}