package main

import "fmt"

// ฟังก์ชันหาลำดับ Fibonacci โดยใช้ Dynamic Programming
func fibonacciDP(n int) int {
	if n <= 1 {
		return n
	}
	fib := make([]int, n+1)
	fib[0], fib[1] = 0, 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}

func main() {
	n := 10 // ต้องการหา Fibonacci ที่ตำแหน่ง n
	fmt.Printf("Fibonacci of %d = %d\n", n, fibonacciDP(n))
}
