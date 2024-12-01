package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func randomSequence(num int) []int {
	var sequence []int
	sequence = make([]int, num, num)
	rand.Seed(uint64(time.Now().UnixNano()))
	var i int
	for i = 0; i < num; i++ {
		sequence[i] = rand.Intn(999) - rand.Intn(999)
	}
	return sequence
}

func InsertionSorter(elements []int) {
	var n = len(elements)
	var i int
	for i = 1; i < n; i++ {
		var j int
		j = i
		for j > 0 {
			if elements[j-1] > elements[j] {
				elements[j-1], elements[j] = elements[j], elements[j-1]
			}
			j = j - 1
		}
	}
}
func main() {
	var sequence []int
	sequence = randomSequence(24)
	fmt.Println("\n^^^^^^^ Before Sorting ^^^ \n\n", sequence)
	InsertionSorter(sequence)
	fmt.Println("\n--- After Sorting ---\n\n", sequence, "\n")
}
