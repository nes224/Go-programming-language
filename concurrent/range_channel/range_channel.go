package main

import (
	"fmt"
	"math/rand"
	"time"
)

func readFromChannel(channel chan int) {
	for number := range channel {
		fmt.Printf("Got '%d' from channel\n", number)
	}
}

func main() {
	channel := make(chan int)
	go readFromChannel(channel)

	for true {
		time.Sleep(1 * time.Second)
		channel <- rand.Intn(100)
	}
}
