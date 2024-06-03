package main

import "fmt"

// การที่ method เรียกใช้ interface จะเรียกว่าเป็นการ implement interface

type Number interface {
	int | float64
}

/*
	"signature" ของ method หรือ function หมายถึงการระบุโครงสร้างของ method หรือ function นั้นๆ ว่าประกอบด้วยอะไรบ้าง
	ซึงรวมถึง
	1. ชื่อของ method หรือ function
	2. ชนิดของ parameter ที่รับเข้า
	3. ชนิดของค่าที่คืนกลับ
*/

func PrintNumber[V Number](x V) {
	fmt.Println(x)
}

func main() {
	PrintNumber(30)
	PrintNumber(30.5)
}
