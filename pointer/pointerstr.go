package main

import "fmt"

type Product struct {
	name string
	price int
	stock int
}

func main() {
	// book := Product{"Golang Book", 250, 50}
	// ptr := &book
	// fmt.Println(ptr)
	// fmt.Println("Product name : ", ptr.name)
	// fmt.Println("Product price : ", ptr.price)
	// fmt.Println("Product stock : ", ptr.stock)
	ptr := &Product{name: "Golang Book", price: 250, stock: 50}
	fmt.Println(ptr)
}