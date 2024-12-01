package main

import (
	"fmt"
)

func add(matrix1 [2][2]int, matrix2 [2][2]int) [2][2] int {
	var m int
	var l int
	var sum [2][2]int
	for l = 0;l < 2;l++{
		for m = 0;m<2;m++{
			sum[l][m] = matrix1[l][m] + matrix2[l][m]
		}
	}
	return sum
}

func subtract(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var m int
	var l int
	var difference [2][2]int
	for l = 0; l < 2; l++ {
		for m =0; m < 2; m++ {
			difference[l][m] = matrix1[l][m] - matrix2[l][m]
		}
	}
	return difference
}

func multiply(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var m int
	var l int
	var n int
	var product [2][2]int
	for l = 0; l < 2; l++ {
		for m = 0; m < 2; m++ {
			var productSum int = 0
			for n =0 ;n<2; n++ {
				for n = 0;n<2;n++ {
					productSum = productSum + matrix1[l][n]*matrix2[n][m]
				}
				product[l][m] = productSum;
			}
		}
	}
	return product
}

func transpose(matrix1 [2][2]int) [2][2]int {
	var m int 
	var l int
	var transMatrix [2][2]int
	for l = 0; l<2;l++{
		for m = 0;m<2;m++ {
			transMatrix[l][m] = matrix1[m][l]
		}
	}
	return transMatrix
}

func determinant(matrix1 [2][2]int) float32 {
	var det float32
	det = det + ((float32(matrix1[0][0]*matrix1[1][1]) - (float32(matrix1[0][1]*matrix1[1][0]))))	
	return det
}

// func inverse(matrix [2][2]int) [][]float64 {
// 	var det float64
// 	det = float64(determinant(matrix))
// 	var invmatrix [][]float64
// 	invmatrix[0][0] = matrix[1][1]/int(det)
// 	invmatrix[0][1] = -1*matrix[0][1]/int(det)
// 	invmatrix[1][0] = -1*matrix[1][0]/int(det)
// 	invmatrix[1][1] = matrix[0][0]/int(det)
// 	return invmatrix
// }

func main() {
	var sum [2][2]int
	matrix1 := [2][2]int{{1,2},{2,1}}
	matrix2 := [2][2]int{{2,1},{1,2}}
	sum = add(matrix1,matrix2)
	fmt.Println("sum:", sum)

	var difference [2][2]int
	difference = subtract(matrix1,matrix2)
	fmt.Println("difference:", difference)

	var product [2][2]int
	product = multiply(matrix1,matrix2)
	fmt.Println("multiply:", product)
}