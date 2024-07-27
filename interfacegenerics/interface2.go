package main 

import "fmt"

type ProductDetails interface {
	details()
}

type ProductStock interface {
	checkStock()
}

type product struct {
	name string
	category string
	price int
	stock int
}

func (d product) details() {
	fmt.Println("Product Name : ", d.name)
	fmt.Println("Product Category : ", d.category)
	fmt.Println("Product Price : ", d.price)
}

func (s product) checkStock() {
	if s.stock <= 10 {
		fmt.Printf("There are %d, almost out of stock.", s.stock)
	} else {
		fmt.Printf("There are %d in stack.", s.stock)
	}
}

func main() {
	p := product{"xPad 10", "Tablet", 18000, 8}
	var pd ProductDetails = p
	pd.details()
	var ps ProductStock = p
	ps.checkStock()
}