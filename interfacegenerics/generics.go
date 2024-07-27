// Generics เป็น templete หรือ format ที่ทำให้เราสามารถนำ code
// ที่เขียนด้วย Golang ใช้งานกับชนิดข้อมูลที่ต่างกันได้ เพื่อช่วยลดความผิดพลาดในการระบุชนิดข้อมูลที่ไม่ตรงกัน
package main

import (
	"fmt"
	// golang.org/x/exp/constraints
)

type Ordered interface {
	int | int32 | int64 | float32 | float64 | string
}

func Max[T Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println("Go Generics Tutorial")
	var a, b int64 = 14, 6
	fmt.Println(Max(a, b))
}
