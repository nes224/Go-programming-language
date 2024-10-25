package main

import "fmt"

func main() {
	var number int
	fmt.Print("กรุณากรอกข้อมูล: ")
	fmt.Scan(&number)

	if number <= 1 {
		fmt.Println("เลขที่ใส่ต้องมากกว่า 1")
		return
	}

	isComposite := false
	arr := make([]int,0)
	for i := 2; i < number; i++ {
		if number%i == 0 {
			arr = append(arr, i)
			isComposite = true
		}
	}
	if !isComposite {
		fmt.Printf("\n%d เป็นจำนวนเฉพาะ (Prime Number)\n", number)
	} else {
		fmt.Printf("\nตัวเลขที่สามารถหารลงตัวได้คือ: %d", number)
		fmt.Printf("\nCan be divisible by %v ", arr)
		fmt.Printf("\n%d เป็นจำนวนประกอบ (Composite Number)\n", number)
	}
}