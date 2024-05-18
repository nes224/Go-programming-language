package main

import "fmt"

func main() {
	c := make(chan bool)
	<-c
	fmt.Println("Done")
}
