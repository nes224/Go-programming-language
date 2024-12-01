package main

import "fmt"

func bubbleSorter(intergers [11]int) {
	var num int
	num = 11
	var isSwapper bool
	isSwapper = true
	for isSwapper {
		isSwapper = false
		var i int
		for i = 1; i < num; i++ {
			if intergers[i-1] > intergers[i] {
				var temp = intergers[i]
				intergers[i] = intergers[i-1]
				intergers[i-1] = temp
				isSwapper = true
			}
		}
	}
	fmt.Println(intergers)
}

func main() {
	var intergers [11]int = [11]int{31,13,12,4,18,16,7,2,3,0,10}
	fmt.Println("Buble Sorter")
	bubbleSorter(intergers)
}
