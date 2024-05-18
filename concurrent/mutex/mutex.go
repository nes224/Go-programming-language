package main

import (
	"fmt"
	"sync"
	"time"
)

// mutex -> Mutual execution is concept of share resource between thread but it mustn't do a tasks at the same time.
// that because we don't conflict accur like a Race condition.
// Golang come with mutex, we can call sync.Mutex and we can use it by call a 2 methods lock and unlock

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock os only one gourtine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	// map key and value
	// ['a'] = 7
	// key = a && value = 7
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))

}
