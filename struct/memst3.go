package main

import "fmt"

type Sales struct {
	name   string
	price  int
	amount int
}

func main() {
	s := Sales{"CameraVlog", 26000, 2}
	display(s)
}

func display(r Sales) {
	fmt.Println("Name:", r.name)
	fmt.Println("Price:", r.price)
	fmt.Println("Amount:", r.amount)
}
