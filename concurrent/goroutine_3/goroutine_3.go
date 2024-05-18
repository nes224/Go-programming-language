package main

import (
	"fmt"
	"time"
)

func goroutine1(c chan bool){
	fmt.Println("Goroutines #1 has started, waiting for Goroutines #2 to start")
	<- c // read from channel
	fmt.Println("Goroutines #1 received a notification from Gourtines #2")
}

func goroutine2(c chan bool) {
	fmt.Println("Goroutines #2 has started, do some work and notify Gorutines #1")
	time.Sleep(2 * time.Second)
	c <- true  // write to channel
	fmt.Println("Goroutine #2 has finished")
}

func main() {
	c := make(chan bool)

	go goroutine1(c)
	go goroutine2(c)

	time.Sleep(3 * time.Second)

}
