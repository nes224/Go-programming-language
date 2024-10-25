package main

import (
	"fmt"
	"math"
)

func main() {
	matrixA := [][]float64{
		{5, 2, 0, 0},
		{2, 5, 2, 0},
		{0, 2, 5, 2},
		{0, 0, 2, 5},
	}
	matrixB := []float64{12, 17, 14, 7}

	// จำนวนของตัวแปร x
	x := make([]float64, len(matrixB))
	tempX := make([]float64, len(matrixB))

	// กำหนดค่า error และตัวแปรควบคุม
	error := 0.000001
	check := true

	for check {
		for i := range x {
			sum := 0.0
			for j := range x {
				if i != j {
					sum += matrixA[i][j] * x[j]
				}
			}
			x[i] = (matrixB[i] - sum) / matrixA[i][i]
		}

		check = false
		for i := range x {
			if math.Abs((x[i]-tempX[i])/x[i]) < error {
				check = false
			} else {
				check = true
			}
			tempX[i] = x[i]
		}
	}

	// แสดงผลลัพธ์
	for i, result := range x {
		fmt.Printf("x[%d] = %.6f\n", i, result)
	}
}
