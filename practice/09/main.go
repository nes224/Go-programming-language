package main

import (
	"fmt"
	"strings"
)

func main() {
	var lines int
	fmt.Print("Enter number of lines: ")
	fmt.Scan(&lines)

	// สร้าง slice สำหรับเก็บค่าของแต่ละบรรทัดในสามเหลี่ยมปาสคาล
	triangle := make([][]int, lines)

	for i := range triangle {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1 // ค่าที่ตำแหน่งแรกในแต่ละบรรทัดคือ 1
		triangle[i][i] = 1 // ค่าที่ตำแหน่งสุดท้ายในแต่ละบรรทัดคือ 1

		// คำนวณค่าในสามเหลี่ยมปาสคาล
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	// แสดงผลรูปสามเหลี่ยมปาสคาล
	for i := 0; i < lines; i++ {
		// พิมพ์ช่องว่างเพื่อจัดตำแหน่งตัวเลขให้อยู่กลาง
		spaces := strings.Repeat("  ", lines-i-1)
		fmt.Print(spaces)

		// พิมพ์ค่าของแต่ละบรรทัด
		for j := 0; j <= i; j++ {
			fmt.Printf("%4d", triangle[i][j]) // พิมพ์ตัวเลขด้วยระยะห่างที่เท่ากัน
		}

		// ขึ้นบรรทัดใหม่
		fmt.Println()
	}
}
