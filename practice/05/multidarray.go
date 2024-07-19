package main

import (
	"fmt"
	"math/rand"
)

func multidarray() {
	var threeArray [2][2][2]int

	var i int
	var j int
	var k int
	for i = 0; i < 2; i++ {
		for j = 0; j < 2; j++ {
			for k = 0; k < 2; k++ {
				threeArray[i][j][k] = rand.Intn(3)
			}
		}
	}
	fmt.Println(threeArray)
}
