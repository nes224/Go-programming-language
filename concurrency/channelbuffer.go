// Buffer หน่วยความจำที่ใช้พักข้อมูลไว้ชั่วคราวในระหว่างที่มีการส่งข้อมูลจำนวน เช่น การโอนไฟล์จำนวนมาก
// เก็บไว้ใน hard disk, Buffered Channel คือ channel ที่มี buffer ไว้รองรับการส่งข้อมูลจำนวนมาก
// ch := make(chan type, buffer), ch := make(chan int, 10)

package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "I love golang."
	ch <- "So good."
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
