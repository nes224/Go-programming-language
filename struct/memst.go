package main

import "fmt"

type Product struct {
	id    int
	name  string
	price int
}

func main() {
	var pd1 Product
	var pd2 = new(Product)
	pd1.id = 135
	pd1.name = "Photo Book"
	pd1.price = 225
	pd2.id = 136
	pd2.name = "Calculator"
	pd2.price = 300

	fmt.Println("Product1")
	fmt.Println("ID:", pd1.id, "Name:", pd1.name, "Price:", pd1.price)
	fmt.Println("Product2")
	fmt.Println("ID:", pd2.id, "Name:", pd2.name, "Price:", pd2.price)
}
