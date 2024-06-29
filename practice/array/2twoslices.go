package main

import "fmt"

func twodslices() {
	var rows int
	var cols int

	rows = 7
	cols = 9
	var twodslice = make([][]int, rows)
	for i := range twodslice {
		twodslice[i] = make([]int, cols)
	}
	fmt.Println("rows:",len(twodslice))
	fmt.Println("cols:",len(twodslice[0]))
	fmt.Println(twodslice)
}
