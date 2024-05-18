package main

import (
	"fmt"
	"time"
)

// optimum point
// overhead
// error out of memory

func main() {
	c := make(chan bool, 3)

	for true {
		go func() {
			fmt.Println(time.Now().Second())
			time.Sleep(2 * time.Second)
			<- c // read data from channel and release the memory
		}()

		c <- true
	}
}
