// Range และ Close, Close -> เป็น function ที่ฝั่งผู้ส่งระบุว่าได้ทำการปิด channel เพราะไม่มีการส่งค่าใด ๆ ไปยัง channel อีกแล้ว
// variable, status := <- ch, Range -> ใช้ร่วมกับ for ในการลูปเพื่อรับค่าจาก channel ไปเรื่อยๆ จนกว่าจะมีการ close channel
// for val := range ch

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	go send(ch)
	time.Sleep(2 * time.Second)
	for val := range ch {
		fmt.Println("Read from ch : ", val)
		time.Sleep(2 * time.Second)
	}
}

func send(ch chan int) {
	for i := 0; i < 4; i++ {
		ch <- i
		fmt.Println("Sent", i, "-> ch : success")
	}
	close(ch)
}