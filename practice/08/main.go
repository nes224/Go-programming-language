package main

import "fmt"

func main() {
	rows := 5

	for i := 0; i < rows; i++ {
		fmt.Print("#")
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		fmt.Print("#")
		fmt.Println()
	}
}