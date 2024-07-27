package main

import "fmt"

type ProductDes interface {
	display()
}

type ProductStock interface {
	ProductDes // Embedded interfaces
	stockValue()
}
// encapsulation private lowercase is private
type product struct {
	name string
	category string
	price int 
	stock int
}

func (d product) display() {
	fmt.Println("Product Name : ", d.name)
	fmt.Println("Product Category : ", d.category)
	fmt.Println("Product Price : ", d.price)
	fmt.Println("Product Stock : ", d.stock)
}

func (s product) stockValue() {
	sv := s.price * s.stock
	fmt.Printf("%s has Stock Value : %d\n", s.name, sv)
}

func main() {
	var pd ProductStock = product{"xPad 10", "Tablet", 18000, 8}
	pd.display()
	pd.stockValue()
}