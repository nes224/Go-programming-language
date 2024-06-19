package main

type Queue []*Order

type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

func (order *Order) New(priority int, quantity int, product string, customerName string) {

}

func main() {

	
}