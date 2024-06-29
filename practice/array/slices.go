package main

import "fmt"

func twiceValue(slice []int) {
	for i, value := range slice {
		slice[i] = 2 * value
	}
}

func slices() {
	slice := []int{1, 3, 5, 6}
	twiceValue(slice)

	for _, value := range slice {
		fmt.Printf("%d\n", value)
	}
}
