package main

import (
	"fmt"
	"strings"
)

func isVowels(c rune) bool {
	return strings.ContainsRune("aeiouAEIOU", c)
}

func main() {
	// rune = int32
	inputString := "ewjqieofiodsacxzmfdksqwpoewqok,;a;csadjkjk3214329ejqwoid"
	size := 5
	vowelsCount := make(map[rune]int)
	var maxVowel rune
	maxCount := 0

	for i := 0; i < len(inputString); i++ {
		end := i + size
		if end > len(inputString) {
			end = len(inputString)
		}
		subStr := inputString[i:end]
		for _, c := range subStr {
			if isVowels(c) {
				vowelsCount[c]++
				if vowelsCount[c] > maxCount {
					maxCount = vowelsCount[c]
					maxVowel = c
				}
			}
		}
	}

	fmt.Println("Vowel counts:", vowelsCount)
	fmt.Printf("The vowel with highest count is '%c'  with %d occurrences\n", maxVowel, maxCount)
}
