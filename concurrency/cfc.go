package main

import (
	"context"
	"fmt"
	"time"
)

// I can’t stress this enough: managing goroutine lifecycles is crucial.
// The context package is your friend here. It lets you cancel operations gracefully,
// which is a big deal in long-running applications.

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		// Simulate some work
		time.Sleep(2 * time.Second)
		cancel() // Cancel the context after 2 seconds
	}()
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Operation complated")
	case <-ctx.Done():
		fmt.Println("Operation canceled")
	}
}
