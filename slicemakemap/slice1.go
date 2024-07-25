package main

import "fmt"

func main() {
	var numbers []int

	numbers = make([]int, 5)
	for i := 0; i< len(numbers); i++ {
		numbers[i] = i*i
	}

	fmt.Println(numbers)
}