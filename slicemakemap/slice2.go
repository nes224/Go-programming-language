package main

import "fmt"

func main() {
	matrix := make([][]int, 3*5)
	fmt.Println("Slice matrix data")
	rows := 3
	cols := 5
	for i := 0; i<rows;i++ {
		matrix[i] = make([]int, 5)
		for j := 0;j<cols;j++ {
			matrix[i][j] = i + j
		}
	}
	fmt.Println(matrix)

}