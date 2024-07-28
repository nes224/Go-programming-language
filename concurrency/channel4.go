package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "receive from server1"
}

func server2(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "receive from server2"
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go server1(c1)
	go server2(c2)
	select {
	case msg1 := <-c1:
		fmt.Println(msg1)
	case msg2 := <-c2:
		fmt.Println(msg2)
	}
}
