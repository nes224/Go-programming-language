package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollDice() int {
	return rand.Intn(6) + 1
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var balance int
	fmt.Print("กรุณากรอกข้อมูล: ")
	fmt.Scan(&balance)

	for balance > 0 {
		dice1 := rollDice()
		dice2 := rollDice()
		sum := dice1 + dice2

		fmt.Printf("คุณทอยได้ลูกเต๋า: %d และ %d ผลรวมคือ: %d\n", dice1, dice2, sum)

		if sum == 7 {
			balance +=4
			fmt.Println("คุณชนะ! คุณได้เงิน 4 บาท")
		} else {
			balance -=1
			fmt.Println("คุณแพ้! คุณเสียเงิน 1 บาท")
		}

		fmt.Printf("เงินคงเหลือ: %d บาท\n", balance)

		if balance > 0 {
			var choice string
			fmt.Print("ต้องการเล่นต่อหรือไม่? (y/n): ")
			fmt.Scan(&choice)
			if choice != "y" {
				break
			}
		} else {
			fmt.Println("เงินหมดแล้ว! เกมจบ")
		}
	}
	fmt.Println("เกมจบ ขอบคุณที่เล่น!")
}
