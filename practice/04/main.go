package main

import (
	"fmt"
	"strconv"
)


func main() {
	// rune หมายถึงประเภทข้อมูลที่ใช้ในการแทนค่าอักขระ (character) ที่เป็น Unicode
	// ซึ่ง rune เป็นชนิดข้อมูลที่เป็น alias ของ int32 ใน Go
	var inputString string
	fmt.Printf("กรุณากรอกข้อมูล: ")
	fmt.Scan(&inputString)

	sum := 0
	for _, v := range inputString {
		num, err := strconv.Atoi(string(v))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		sum += num
	}

	fmt.Printf("ผลรวมของตัวเลข %s คือ %d\n", inputString, sum)
}