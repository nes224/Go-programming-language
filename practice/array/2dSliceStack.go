package main

import "fmt"

func twodSliceStack() {
	originNumber := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rows := 3
	cols := 4
	twodSlice := make([][]int, rows)

	for i := 0; i < rows; i++ {
		twodSlice[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			index := i*cols + j
			if index < len(originNumber) {
				twodSlice[i][j] = originNumber[index]
			}
		}
	}

	fmt.Println(twodSlice)
}

