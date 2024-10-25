package main

import "fmt"

func main() {
	var n int
	fmt.Print("กรุณาใส่จำนวนลำดับที่ต้องการหาผลรวม: ")
	fmt.Scan(&n)

	sum := 0.0
	numerator := 1
	for i := 0; i < n; i++ {
		denominator := numerator + 1
		sum += float64(numerator) / float64(denominator)
		numerator += 2
	}

	fmt.Printf("ผลรวมของ %d ลำดับแรกคือ: %.6f\n", n, sum)
}
