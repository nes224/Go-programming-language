package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan bool)
	rand.Seed(time.Now().Unix())

	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		channel <- true
	}()

	timeout := time.Tick(3 * time.Second)
	select {
	case <-timeout:
		fmt.Println("Timeout!")
	case <- channel:
		fmt.Println("Work done!")
	}
}
