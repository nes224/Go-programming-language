package main

import "fmt"

func main() {
	var number int
	fmt.Print("กรุณาใส่ตัวเลขจำนวนเต็ม:")
	fmt.Scan(&number)

	if number%5 == 0 && number%6 == 0 {
		fmt.Printf("%d หารด้วย 5 และ 6 ลงตัว \n", number)
	} else if number%5 == 0 || number %6 == 0 {
		fmt.Printf("%d หารด้วย 5 หรือ 6 ลงตัว แต่ไม่ทั้งสอง\n", number)
	} else {
		fmt.Printf("%d ไม่สามารถหารด้วย 5 หรือ 6 ลงตัว\n", number)
	}
}
