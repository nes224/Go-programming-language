package main

import (
	"fmt"
	"time"
)

func main() {
	go HelloOne()
	go HelloTwo()
	time.Sleep(3600 * time.Millisecond)
	fmt.Println("main function")
}

func different() int64 {
	return time.Now().UnixMicro() / int64(time.Microsecond)
}

func HelloOne() {
	for i := 0; i <= 5; i++ {
		start := different()
		time.Sleep(250 * time.Millisecond)
		end := different()
		fmt.Printf("No.%d %v Duration (ms) : %d\n", i, "Run HelloOne", end-start)
	}
}

func HelloTwo() {
	for i := 0; i <= 5; i++ {
		start := different()
		time.Sleep(400 * time.Millisecond)
		end := different()
		fmt.Printf("No.%d %v Duration (ms) : %d\n", i, "Run HelloTwo", end-start)
	}
}

// HelloOne 0 -> 250 -> 500 -> 750 -> 1000 -> 1250 (ms)
// HelloTwo 0 -----> 400 -------> 800 ------> 1200 --------> 1600 ------> 2000 (ms)
// Main ----------------------------------------------------------------------------------------> 3000 (ms)
