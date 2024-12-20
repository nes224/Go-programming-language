package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func createArray(num int) []int {
	var array []int
	array = make([]int, num, num)
	rand.Seed(uint64(time.Now().UnixNano()))
	var i int
	for i = 0; i < num; i++ {
		array[i] = rand.Intn(99999) - rand.Intn(99999)
	}
	return array
}

func MergeSorter(array []int) []int {
	if len(array) < 2 {
		return array
	}
	var middle int
	middle = (len(array)) / 2
	return JoinArrays(MergeSorter(array[:middle]), MergeSorter(array[middle:]))
}

func JoinArrays(leftArr []int, rightArr []int) []int {
	var num int
	var i int
	var j int
	num, i, j = len(leftArr)+len(rightArr), 0, 0
	var array []int
	array = make([]int, num, num)
	var k int
	for k = 0; k < num; k++ {
		if i > len(leftArr)-1 && j <= len(rightArr)-1 {
			array[k] = rightArr[j]
			j++
		} else if j > len(rightArr)-1 && i <= len(leftArr)-1 {
			array[k] = leftArr[i]
			i++
		} else if leftArr[i] < rightArr[j] {
			array[k] = leftArr[i]
			i++
		} else {
			array[k] = array[j]
			j++
		}
	}
	return array
}

func main() {
	var elements []int
	elements = createArray(40)
	fmt.Println("\n Before Sorting \n\n", elements)
	fmt.Println("\n-After Sorting\n\n", MergeSorter(elements), "\n")
}
