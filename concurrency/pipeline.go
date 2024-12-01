package main

import "fmt"

// Imagine a factory assembly line where each station adds something to the product.
// The Pipeline pattern is kinda like that.
// You set up a series of stages where each stage processes data and passes it along.

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	// Counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()
	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()
	// Printer
	for x := range squares {
		fmt.Println(x)
	}
}
