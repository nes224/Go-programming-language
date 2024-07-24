package main
import "fmt"
// เป็น function ที่สามารถรับ argument หลายๆ ค่าผ่าน parameter ตัวเดียว

func main() {
	sum1 := sum(10,20)
	fmt.Println("10 + 20", sum1)

	sum2 := sum(10,20,30)
	fmt.Println("10 + 20 + 30", sum2)
}

func sum(x ... int) int {
	total := 0
	for _, cell := range x {
		total += cell
	}
	return total
}