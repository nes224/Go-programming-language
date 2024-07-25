package main 

import "fmt"

type Rectangle struct {
	width, length int
}

func (r *Rectangle) scale(m int) {
	r.width = r.width * m
	r.length = r.length * m
}

func (r *Rectangle) Area() int {
	return r.width * r.length
}

func main() {
	rec := &Rectangle{10, 5}
	fmt.Printf("Before scale width : %d, length : %d \n", rec.width, rec.length)
	fmt.Printf("Area before scale : %d\n", rec.Area())
	rec.scale(2)
	fmt.Printf("After scale width : %d, length : %d \n", rec.width, rec.length)
	fmt.Printf("Area After scale : %d\n", rec.Area())
}