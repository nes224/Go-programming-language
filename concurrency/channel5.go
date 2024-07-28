package main 

import (
	"fmt"
	"time"
)

func server(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Receive from server"
}

func main() {
	c := make(chan string)
	go server(c)

	for {
		time.Sleep(300 * time.Millisecond)
		select {
		case msg := <- c:
			fmt.Println(msg)
			return
		default:
			fmt.Println("No message from server")
		}
	}
}