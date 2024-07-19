package main

import "fmt"

func add(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var m int
	var l int
	var sum [2][2]int
	for l = 0; l < 2; l++ {
		for m = 0; m < 2; m++ {
			sum[l][m] = matrix1[l][m] + matrix2[l][m]
		}
	}
	return sum
}

func twodmatrix() {
	var matrix1 = [2][2]int{
		{4, 5},
		{1, 2}}
	var matrix2 = [2][2]int{
		{6, 7},
		{3, 4},
	}

	sum := add(matrix1, matrix2)
	fmt.Println(sum)

}
