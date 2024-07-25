package main

import "fmt"

func main() {
	arr := [5]int{11,22,33,44,55}
	var ptr *[5]int
	fmt.Println("Before :",arr)
	ptr = &arr
	(*ptr)[0] = 1

	fmt.Println("After :",arr)
}