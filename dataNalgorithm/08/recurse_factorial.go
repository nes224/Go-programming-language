package main

import (
	"fmt"
)

func Factor(num int) int {
	if num <= 1 {
		return 1
	}

	return num * Factor(num-1)
}

func main() {
	var num int = 12 
	fmt.Println("Factorial: %d is %d", num, Factor(num))
}