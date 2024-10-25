package main

import (
	"fmt"
	"strings"
)

func main() {
	inputString := "test go with you don't do stupid things you need to calm down and keep strong on yourself and believe in yourself."
	countA, countE, countI, countO, countU := 0, 0, 0, 0, 0
	line := strings.ToLower(inputString)
	for _, char := range line {
		switch char {
		case 'a':
			countA++
		case 'e':
			countE++
		case 'i':
			countI++
		case 'o':
			countO++
		case 'u':
			countU++
		}
	}

	fmt.Printf("จำนวนตัวอักษร a: %d\n", countA)
	fmt.Printf("จำนวนตัวอักษร e: %d\n", countE)
	fmt.Printf("จำนวนตัวอักษร i: %d\n", countI)
	fmt.Printf("จำนวนตัวอักษร o: %d\n", countO)
	fmt.Printf("จำนวนตัวอักษร u: %d\n", countU)

	maxVowel := countA
	maxChar := 'a'

	if countE > maxVowel {
		maxVowel = countE
		maxChar = 'e'
	}

	if countI > maxVowel {
		maxVowel = countE
		maxChar = 'i'
	}

	if countO > maxVowel {
		maxVowel = countO
		maxChar = 'o'
	}

	if countU > maxVowel {
		maxVowel = countU
		maxChar = 'u'
	}

	fmt.Printf("สระที่มีจำนวนมากที่สุดคือ %c มีจำนวน %d ตัว\n", maxChar, maxVowel)
}
