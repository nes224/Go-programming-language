package main

import "fmt"

// Connection Stack Database
func main() {
	// defer -> Stack or Last in first out
	defer fmt.Println("defer: return_1")
	defer fmt.Println("defer: return_2")
	tests := []int{1, 2, 3, 4, 5}
	for _, test := range tests {
		fmt.Printf("%d", test)
	}
	fmt.Println()
}
