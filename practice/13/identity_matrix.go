package main

import (
	"fmt"
)

func Matrix(order int) [][]float64 {
	var matrix [][]float64
	matrix = make([][]float64, order)
	var i int
	for i = 0; i < order; i++ {
		var temp []float64
		temp = make([]float64, order)
		temp[i] = 1 
		matrix[i] = temp
	}
	return matrix
}

func identity() {
	fmt.Println(Matrix(4))
}