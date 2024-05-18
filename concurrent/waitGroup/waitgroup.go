package main

import (
	"fmt"
	"sync"
	"time"
)

// go routines
// WaitGroup is blocking function main and waiting all goroutines done.
func main() {
	wg := sync.WaitGroup{}

	// add: increase goroutines
	wg.Add(1)
	go doWork(&wg, "A")

	wg.Add(1)
	go doWork(&wg, "B")

	wg.Add(1)
	go doWork(&wg, "C")

	// block goroutine until it become zero
	wg.Wait()
	fmt.Println("All done!")
}

func doWork(wg *sync.WaitGroup, name string) {
	fmt.Printf("doWork for %s\n", name)
	time.Sleep(2 * time.Second)
	// done: decrease goroutines
	wg.Done()
}
