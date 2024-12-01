package main

import (
	"fmt"
)

func changeMatrix(matrix [3][3]int) [3][3]int {
	var i int 
	var j int 
	var Rows [3]int
	var Columns [3]int 

	var matrixChanged [3][3]int 
	for i =0; i < 3; i++ {
		for j=0; j < 3; j++ {
			if matrix[i][j] == 1 {
				Rows[i] = 1 
				Columns[j] = 1  
			}
		}
	}
	for i = 0;i<3;i++{
		for j =0;j<3;j++{
			if Rows[1] == 1 || Columns[j] == 1 {
				matrixChanged[i][j] = 1 
			}
		}
	}
	return matrixChanged
}

func printMatrix(matrix [3][3]int) {
	var i int 
	var j int
	for i = 0; i<3;i++{
		for j =0;j<3;j++{
			fmt.Printf("%d", matrix[i][j])
		}
		fmt.Printf("\n")
	} 
}

func main() {
	var matrix = [3][3]int{{1,0,0},{0,0,0},{0,0,0}}
	printMatrix(matrix)
	matrix = changeMatrix(matrix)
	printMatrix(matrix)
}