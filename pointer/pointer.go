package main

import "fmt"

func main() {
	var x int = 5
	var p *int

	fmt.Println("Value of x = ", x , "Address of x = ", &x)
	fmt.Println("Value of p =", p , "Address of p = ", &p)

	p = &x
	*p = 10

	fmt.Println("-----------------------------------------------")
	
	fmt.Println("Value of x = ", x , "Address of x =", &x)
	fmt.Println("Value of *p = ", *p, "Value of p =", p)
}