package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	checkValues := []int{15, 18, 22, 25, 30}

	num1 := rand.Intn(11) + 10
	num2 := rand.Intn(11) + 10

	sum := num1 + num2

	fmt.Printf("สุ่มได้เลข: %d และ %d\n", num1, num2)
	fmt.Printf("ผลรวมของ %d และ %d คือ: %d\n", num1, num2, sum)

	found := false
	for _, v := range checkValues {
		if sum == v {
			found = true
			break
		}
	}
	if found == true {
		fmt.Printf("ผลรวม %d ตรงกับค่าที่มีใน slice\n", sum)
	} else {
		fmt.Printf("ผลรวม %d ไม่ตรงกับค่าที่มีใน slice\n", sum)
	}
}
