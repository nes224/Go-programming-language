package main

// empty interface = data type any

import "fmt"

func display(d interface{}) {
	fmt.Println(d)
}

func main() {
	x := [3]string{"Windows", "MAC", "Linux"}
	display(12)
	display("Tester Nes Golang")
	display(true)
	display(x)
}

// Type assertions