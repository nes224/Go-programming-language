package main

import (
	"fmt"
	"sync"
)

// Every system in scale should have tasks that are merging and diverging asynchronously.
// We can use Fan-Out, Fan-In pattern here,
// like fanout by distributing the load to parallel goroutines and fan in thereby collecting their results

func main() {
	in := gen(2, 3, 4, 5)
	// Fan-out
	ch1 := sq(in)
	ch2 := sq(in)
	// Fan-in
	for n := range merge(ch1, ch2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int {
	var wg sync.WaitGroup
	out := make(chan int) 
	output := func(c chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
