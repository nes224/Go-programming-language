package main

import "fmt"

func hello(name string, result chan<- string) {
	output := "Hello " + name
	fmt.Println("In function = %s\n", output)
	result <- output // ผลลัพธ์การทำงานจะถูกส่งเข้าไปยัง channel ที่ส่งมา
}

func goroutineFunc1() {
	result := make(chan string) // สร้าง channel ชื่อว่า result โดยชนิดข้อมูลที่ส่งใน channel คือ string และทำการส่งเข้าไปใน hello()
	fakeName := "Fake Name"
	go hello(fakeName, result) 
	<-result // main() จะรอผลการทำงานจาก channel ด้วย <-result
	fmt.Println("Finish main")

}
