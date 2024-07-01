package main

import "fmt"

func h(x, y int) int {
	return x * y
}

func g(l, m int) (x, y int) {
	x = 2 * l
	y = 4 * m
	return
}

func Tuples() {
	fmt.Println(h(g(2, 4)))
}
