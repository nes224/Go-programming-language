package main

import (
	"fmt"
	"time"
)

// <- เป็นท่อสำหรับส่งข้อมูล <- คือตัวดำเนินการ สื่อความหมายคือการไหลของข้อมูล
/*
	c := make(chan int) กำหนด c เป็นตัวแปร channel
	c <- 10 ส่งค่า 10 ให้กับ channel c
	x := <- c x รับค่าจาก channel c
 Golang จะ block การส่งและรับข้อมูลที่กำลังดำเนินการอยู่จนกว่าอีกฝั่งจะพร้อมทำงาน
 ได้ประสานพอดีกัน
*/


func main() {
	ch := make(chan int)
	for i := 0; i < 4;i++ {
		go hello(i, ch)
	}

	for i := 0;i < 4 ;i++ {
		id := <- ch
		fmt.Println("Receive value from channel", id)
	}
}

func hello(n int, ch chan int) {
	for i := 0; i < 4; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Send value to channel", n)
		ch <- n
	}
}