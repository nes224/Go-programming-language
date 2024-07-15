package main

import "fmt"

func changeMatrix(matrix [3][3]int) [3][3]int {
	var i int
	var j int
	var Rows [3]int
	var Cols [3]int

	var matrixChange [3][3]int
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if matrix[i][j] == 1 {
				Rows[i] = 1
				Cols[j] = 1
			}
		}
	}

	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			if Rows[i] == 1 || Cols[j] == 1 {
				matrixChange[i][j] = 1
			}
		}
	}
	return matrixChange
}

func printMatrix(matrix [3][3]int) {
	var i int
	var j int
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			fmt.Printf("%d", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}

func boolMatrix() {
	matrix := [3][3]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	printMatrix(matrix)
	matrix = changeMatrix(matrix)
	printMatrix(matrix)
}
